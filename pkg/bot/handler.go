package bot

import (
	"context"
	"echo/internal/biz"
	"echo/internal/model"
	"echo/internal/registry"
	"echo/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/lxzan/gws"
	"sync"
	"time"
)

type Handler struct {
	gws.BuiltinEventHandler
	conns       map[*gws.Conn]struct{}
	mu          sync.Mutex
	registry    *registry.CommandRegistry
	logger      *logger.Logger
	bot         *biz.BotUsecase
	apiHandlers map[*gws.Conn][]struct {
		handler func(*model.OneBotMessage, map[string]interface{}) string
		msg     *model.OneBotMessage
		sentAt  time.Time
	}
	sendHandlers map[*gws.Conn][]struct {
		msg    *model.OneBotMessage
		conn   *gws.Conn
		sentAt time.Time
	}
}

func NewBotHandler(registry *registry.CommandRegistry, log *logger.Logger, bot *biz.BotUsecase) *Handler {
	return &Handler{
		registry: registry,
		logger:   log,
		bot:      bot,
		conns:    make(map[*gws.Conn]struct{}),
		apiHandlers: make(map[*gws.Conn][]struct {
			handler func(*model.OneBotMessage, map[string]interface{}) string
			msg     *model.OneBotMessage
			sentAt  time.Time
		}),
		sendHandlers: make(map[*gws.Conn][]struct {
			msg    *model.OneBotMessage
			conn   *gws.Conn
			sentAt time.Time
		}),
	}
}

func (h *Handler) OnOpen(socket *gws.Conn) {
	h.logger.Info().Str("address", socket.RemoteAddr().String()).Msg("连接建立成功")
	h.mu.Lock()
	h.conns[socket] = struct{}{}
	socket.Session().Store("conn", socket)
	h.mu.Unlock()
}

func (h *Handler) OnClose(socket *gws.Conn) {
	h.logger.Info().Str("address", socket.RemoteAddr().String()).Msg("连接关闭")
	h.mu.Lock()
	delete(h.conns, socket)
	h.mu.Unlock()
}

func (h *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	var data map[string]interface{}
	if err := json.Unmarshal(message.Bytes(), &data); err != nil {
		h.logger.Error().Err(err).Str("message", string(message.Bytes())).Msg("解析消息失败")
		return
	}

	// 先检查是否是 API 响应（无 post_type，有 status 和 retcode）
	if _, hasPostType := data["post_type"]; !hasPostType {
		if _, hasStatus := data["status"]; hasStatus {
			if _, hasRetcode := data["retcode"]; hasRetcode {
				h.handleAPIResponse(socket, data)
				return
			}
		}
	}

	// 处理事件消息（有 post_type）
	if postType, ok := data["post_type"]; ok {
		var msg model.OneBotMessage
		if err := json.Unmarshal(message.Bytes(), &msg); err != nil {
			h.logger.Error().
				Str("app", "echo").
				AnErr("err", err).
				Str("message", string(message.Bytes())).
				Msg("解析消息失败")
			return
		}

		switch postType {
		case "meta_event":
			h.metaEvent(&msg, socket)
		case "message":
			h.logger.Info().Str("message", string(message.Bytes())).Msg("收到指令消息")
			reply, ok := h.registry.Execute(&msg)
			if ok {
				h.sendReply(socket, &msg, reply)
			} else {
				h.logger.Info().Str("message", msg.RawMessage).Msg("指令无响应")
			}
		default:
			h.logger.Info().Str("post_type", postType.(string)).Msg("未处理的事件类型")
		}
	}
}

func (h *Handler) SendAPIRequest(socket *gws.Conn, request []byte, msg *model.OneBotMessage, handler func(*model.OneBotMessage, map[string]interface{}) string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.conns[socket]; !ok {
		h.logger.Error().Str("app", "echo").Msg("连接不可用，无法发送 API 请求")
		return
	}

	h.apiHandlers[socket] = append(h.apiHandlers[socket], struct {
		handler func(*model.OneBotMessage, map[string]interface{}) string
		msg     *model.OneBotMessage
		sentAt  time.Time
	}{handler, msg, time.Now()})

	socket.WriteMessage(gws.OpcodeText, request)
	h.logger.Info().Str("request", string(request)).Msg("API 请求已发送")
}

func (h *Handler) sendReply(socket *gws.Conn, msg *model.OneBotMessage, content string) {
	var action model.OneBotAction
	if msg != nil && msg.MessageType == "private" {
		action = model.OneBotAction{
			Action: "send_private_msg",
			Params: model.SendMessageParams{
				UserID:     msg.UserId,
				Message:    content,
				AutoEscape: false,
			},
		}
	} else if msg != nil && msg.MessageType == "group" {
		action = model.OneBotAction{
			Action: "send_group_msg",
			Params: model.SendMessageParams{
				GroupID:    msg.GroupId,
				Message:    content,
				AutoEscape: false,
			},
		}
	} else {
		action = model.OneBotAction{
			Action: "send_msg",
			Params: model.SendMessageParams{
				Message:    content,
				AutoEscape: false,
			},
		}
	}

	actionBytes, _ := json.Marshal(action)
	h.mu.Lock()
	h.sendHandlers[socket] = append(h.sendHandlers[socket], struct {
		msg    *model.OneBotMessage
		conn   *gws.Conn
		sentAt time.Time
	}{msg, socket, time.Now()})
	h.mu.Unlock()

	if _, ok := h.conns[socket]; ok {
		socket.WriteMessage(gws.OpcodeText, actionBytes)
		h.logger.Info().Str("content", content).Msg("消息已发送")
	} else {
		h.logger.Error().Str("app", "echo").Msg("消息发送失败，无可用链接")
	}
}

func (h *Handler) metaEvent(msg *model.OneBotMessage, socket *gws.Conn) {
	ctx := context.Background()

	if msg.MetaEventType == "lifecycle" && msg.SubType == "connect" {
		h.logger.Info().Int64("bot", msg.SelfId).Msg("BOT连接成功")
		err := h.bot.CreateBot(ctx, msg.SelfId, msg.Time, socket.RemoteAddr().String())
		if err != nil {
			h.logger.Error().
				Str("app", "echo").
				AnErr("err", err).
				Msg("创建 BOT 失败")
			return
		}
	} else if msg.MetaEventType == "heartbeat" {
		err := h.bot.UpdateBotUptime(ctx, msg.SelfId, msg.Time, socket.RemoteAddr().String())
		if err != nil {
			h.logger.Error().
				Str("app", "echo").
				AnErr("err", err).
				Msg("更新 BOT 在线时间失败")
			return
		}
		h.logger.Info().Int64("bot", msg.SelfId).Msg("收到心跳包")
	}
}

func (h *Handler) handleAPIResponse(socket *gws.Conn, data map[string]interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	apiQueue, apiOk := h.apiHandlers[socket]
	sendQueue, sendOk := h.sendHandlers[socket]
	if !apiOk && !sendOk {
		h.logger.Error().Str("app", "echo").Msg("无法处理消息")
		return
	}

	// 判断是否是发送消息的结果
	if dataOk, ok := data["data"]; ok {
		if dataMap, ok := dataOk.(map[string]interface{}); ok {
			if _, hasMessageID := dataMap["message_id"]; hasMessageID {
				if sendOk && len(sendQueue) > 0 {
					handler := sendQueue[0]
					h.sendHandlers[socket] = sendQueue[1:]
					h.handleSendResponse(handler.conn, handler.msg, data)
					return
				}
			}
		}
	}

	// 处理其他 API 响应
	if apiOk && len(apiQueue) > 0 {
		handler := apiQueue[0]
		h.apiHandlers[socket] = apiQueue[1:]
		reply := handler.handler(handler.msg, data)
		h.sendReply(socket, handler.msg, reply)
	} else {
		h.logger.Error().Str("app", "echo").Msg("无法处理消息")
	}
}

func (h *Handler) handleSendResponse(socket *gws.Conn, msg *model.OneBotMessage, data map[string]interface{}) {
	if data["status"] == "ok" && data["retcode"].(float64) == 0 {
		if sendData, ok := data["data"].(map[string]interface{}); ok {
			messageID := int64(sendData["message_id"].(float64))
			h.logger.Info().
				Str("app", "echo").
				Int64("message_id", messageID).
				Int64("groupId", msg.GroupId).
				Int64("botId", msg.SelfId).
				Msg("消息发送成功")
			reply := fmt.Sprintf("消息发送成功，Message ID: %d", messageID)
			h.sendReply(socket, msg, reply)
		}
	} else {
		h.logger.Error().
			Str("app", "echo").
			Int64("groupId", msg.GroupId).
			Int64("botId", msg.UserId).
			Msg("消息发送失败")
	}
}
