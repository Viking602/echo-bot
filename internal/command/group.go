package command

import (
	"context"
	"echo/internal/model"
	"echo/internal/service"
)

type GroupCommand struct {
	moyuService *service.GroupMoyuService
}

func NewGroupCommand(
	moyuService *service.GroupMoyuService,
) *GroupCommand {
	return &GroupCommand{
		moyuService: moyuService,
	}
}

func (c *GroupCommand) Register(registry model.Registrar) {
	ctx := context.Background()
	groupCmd := model.NewCommand("group", func(msg *model.OneBotMessage, args []string) string {

		return "group 帮助\n" +
			"group moyu on｜off  开启/关闭 摸鱼人日历\n"
	}, model.ScopeGroup)

	moyuCmd := model.NewCommand("moyu", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请输入参数 group moyu on｜off"
		}
		return c.moyuService.GroupMoyu(ctx, msg.SelfId, msg.GroupId, args[0])
	}, model.ScopeGroup)

	groupCmd.AddChild(moyuCmd)

	registry.Register(groupCmd)
}
