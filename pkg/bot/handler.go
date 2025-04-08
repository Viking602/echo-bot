package bot

import (
	"context"
	"echo/internal/biz"
	"echo/internal/model"
	"echo/internal/registry"
	"echo/pkg/logger"
	"encoding/json"
	"github.com/lxzan/gws"
)

type Handler struct {
	gws.BuiltinEventHandler
	socket   *gws.Conn
	registry *registry.CommandRegistry
	logger   *logger.Logger
	bot      *biz.BotUsecase
}

func NewBotHandler(registry *registry.CommandRegistry, log *logger.Logger, bot *biz.BotUsecase) *Handler {
	return &Handler{
		registry: registry,
		logger:   log,
	}
}

func (h *Handler) OnOpen(socket *gws.Conn) {
	h.logger.Info().Str("address", socket.RemoteAddr().String()).Msg("连接建立成功")
	h.socket = socket
}

func (h *Handler) OnClose(socket *gws.Conn, err error) {
	h.logger.Info().Str("address", socket.RemoteAddr().String()).Msg("连接关闭")
}

func (h *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	var msg model.OneBotMessage
	if err := json.Unmarshal(message.Bytes(), &msg); err != nil {
		h.logger.Info().
			Str("app", "echo").
			AnErr("err", err).
			Msg("Failed to parse message")
		return
	}

	if msg.PostType == "message" {
		reply, ok := h.registry.Execute(&msg)
		if ok {
			h.sendReply(&msg, reply)
		}
	}
}

func (h *Handler) sendReply(msg *model.OneBotMessage, content string) {
	ctx := context.Background()
	h.bot.CreateBot(ctx, msg.SelfId)
	var action model.OneBotAction
	if msg.MessageType == "private" {
		action = model.OneBotAction{
			Action: "send_private_msg",
			Params: model.SendMessageParams{
				UserID:     msg.UserId,
				Message:    content,
				AutoEscape: false,
			},
		}
	} else if msg.MessageType == "group" {
		action = model.OneBotAction{
			Action: "send_group_msg",
			Params: model.SendMessageParams{
				GroupID:    msg.Sender.UserId,
				Message:    content,
				AutoEscape: false,
			},
		}
	}

	actionBytes, _ := json.Marshal(action)
	h.socket.WriteMessage(gws.OpcodeText, actionBytes)
}
