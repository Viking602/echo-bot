package command

import (
	"echo/internal/model"
	"echo/internal/service"
	"strings"
)

type EchoCommand struct {
	echoService *service.EchoService
}

func NewEchoCommand(echoService *service.EchoService) *EchoCommand {
	return &EchoCommand{echoService: echoService}
}

func (c *EchoCommand) Register(registry model.Registrar) {
	echoCmd := model.NewCommand("echo", func(msg *model.OneBotMessage, args []string) string {
		return c.echoService.Echo(msg)
	}, model.ScopePrivate) // 仅私聊

	repeatCmd := model.NewCommand("repeat", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请提供要重复的内容"
		}
		return strings.Repeat(args[0], 2)
	}, model.ScopePrivate) // 仅私聊

	echoCmd.AddChild(repeatCmd)
	registry.Register(echoCmd)
}
