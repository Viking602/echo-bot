package registry

import (
	"echo/internal/model"
	"fmt"
	"strings"
)

type CommandRegistry struct {
	commands map[string]*model.Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]*model.Command),
	}
}

func (r *CommandRegistry) Register(cmd *model.Command) {
	r.commands[cmd.Name] = cmd
}

func (r *CommandRegistry) Execute(msg *model.OneBotMessage) (string, bool) {
	raw := msg.RawMessage
	if !strings.HasPrefix(raw, "/") {
		return "", false
	}

	parts := strings.Fields(raw[1:])
	if len(parts) == 0 {
		return "", false
	}

	cmd, exists := r.commands[parts[0]]
	if !exists {
		return "未知指令", true
	}

	// 检查命令适用范围
	if !r.isScopeValid(cmd, msg.MessageType) {
		if msg.MessageType == "private" {
			return "此指令仅适用于群聊", true
		}
		return "此指令仅适用于私聊", true
	}

	// 解析子命令
	remainingParts := parts[1:]
	for len(remainingParts) > 0 {
		nextCmd, ok := cmd.Children[remainingParts[0]]
		if !ok {
			if cmd.Handler != nil {
				return cmd.Handler(msg, remainingParts), true
			}
			return "缺少子命令或参数", true
		}
		// 检查子命令适用范围
		if !r.isScopeValid(nextCmd, msg.MessageType) {
			if msg.MessageType == "private" {
				return "此子命令仅适用于群聊", true
			}
			return "此子命令仅适用于私聊", true
		}
		cmd = nextCmd
		remainingParts = remainingParts[1:]
	}

	if cmd.Handler != nil {
		return cmd.Handler(msg, remainingParts), true
	}
	return fmt.Sprintf("%s 指令缺少参数", cmd.Name), true
}

// isScopeValid 检查命令是否适用于当前消息类型
func (r *CommandRegistry) isScopeValid(cmd *model.Command, messageType string) bool {
	if cmd.Scope == model.ScopeAll {
		return true
	}
	if messageType == "private" {
		return cmd.Scope&model.ScopePrivate != 0
	}
	if messageType == "group" {
		return cmd.Scope&model.ScopeGroup != 0
	}
	return false
}
