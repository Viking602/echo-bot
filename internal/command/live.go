package command

import (
	"context"
	"echo/internal/model"
	"echo/internal/service"
)

type LiveCommand struct {
	liveService *service.LiveAddService
}

func NewLiveCommand(liveService *service.LiveAddService) *LiveCommand {
	return &LiveCommand{
		liveService: liveService,
	}
}

func (c *LiveCommand) Register(registry model.Registrar) {
	ctx := context.Background()
	liveCmd := model.NewCommand("live", func(msg *model.OneBotMessage, args []string) string {
		return "请使用子命令 live <平台>"
	}, model.ScopeGroup)

	addCmd := model.NewCommand("add", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请输入平台"
		}
		return "请输入平台"
	}, model.ScopeGroup)

	biliCmd := model.NewCommand("bilibili", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请输入房间号"
		}
		add := c.liveService.AddLive(ctx, msg.SelfId, msg.GroupId, args[0])
		return add
	}, model.ScopeGroup)

	liveCmd.AddChild(addCmd)
	addCmd.AddChild(biliCmd)
	registry.Register(liveCmd)

}
