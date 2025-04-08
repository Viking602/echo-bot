package command

import (
	"echo/internal/model"
	"echo/internal/service"
)

type EchoCommand struct {
	echoService *service.EchoService
}

func NewEchoCommand(echoService *service.EchoService) *EchoCommand {
	return &EchoCommand{echoService: echoService}
}

func (c *EchoCommand) Register(registry model.Registrar) {
	registry.Register(model.Command{
		Name: "echo",
		Handler: func(msg *model.OneBotMessage) string {
			return c.echoService.Echo(msg)
		},
	})
}
