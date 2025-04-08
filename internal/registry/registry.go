package registry

import (
	"echo/internal/model"
	"strings"
)

type CommandRegistry struct {
	commands map[string]model.Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]model.Command),
	}
}

func (r *CommandRegistry) Register(cmd model.Command) {
	r.commands[cmd.Name] = cmd
}

func (r *CommandRegistry) Execute(msg *model.OneBotMessage) (string, bool) {
	raw := msg.RawMessage
	if !strings.HasPrefix(raw, "-") {
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

	return cmd.Handler(msg), true
}
