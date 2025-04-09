package command

import (
	"echo/internal/model"
	"echo/internal/service"
)

type SetCommand struct {
	setMaster *service.SetMasterService
}

func NewSetCommand(setMaster *service.SetMasterService) *SetCommand {
	return &SetCommand{setMaster: setMaster}
}

func (c *SetCommand) Register(registry model.Registrar) {
	setCmd := model.NewCommand("set", func(msg *model.OneBotMessage, args []string) string {
		return "请使用子命令 set master <Token>"
	}, model.ScopePrivate)

	setMasterCmd := model.NewCommand("master", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请输入 Token"
		}
		token := args[0]
		master := c.setMaster.SetMaster(token, msg.SelfId, msg.UserId)
		return master
	}, model.ScopePrivate)
	setCmd.AddChild(setMasterCmd)
	registry.Register(setCmd)
}
