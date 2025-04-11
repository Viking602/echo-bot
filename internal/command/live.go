package command

import (
	"context"
	"echo/internal/model"
	"echo/internal/service"
)

type LiveCommand struct {
	biliAddLiveService *service.BiliLiveAddService
	biliDelLiveService *service.BiliLiveDelService
}

func NewLiveCommand(biliAddLiveService *service.BiliLiveAddService, biliDelLiveService *service.BiliLiveDelService) *LiveCommand {
	return &LiveCommand{
		biliAddLiveService: biliAddLiveService,
		biliDelLiveService: biliDelLiveService,
	}
}

func (c *LiveCommand) Register(registry model.Registrar) {
	ctx := context.Background()
	liveCmd := model.NewCommand("live", func(msg *model.OneBotMessage, args []string) string {
		return "live 帮助 \n" +
			"live add <平台> <房间号> 添加订阅\n" +
			"live del <平台> <房间号> 删除订阅\n" +
			"live all 获取所有订阅"
	}, model.ScopeGroup)

	// 添加逻辑
	addCmd := model.NewCommand("add", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请输入平台 live add <平台> <房间号>"
		}
		return "请输入平台 live add <平台> <房间号>"
	}, model.ScopeGroup)

	biliCmd := model.NewCommand("bilibili", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请输入房间号 live add bilibili <房间号>"
		}
		add := c.biliAddLiveService.AddLive(ctx, msg.SelfId, msg.GroupId, args[0])
		return add
	}, model.ScopeGroup)

	// 删除逻辑
	delCmd := model.NewCommand("del", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请输入平台 live del <平台> <房间号>"
		}
		return "请输入平台 live del <平台> <房间号>"
	}, model.ScopeGroup)

	biliDelCmd := model.NewCommand("bilibili", func(msg *model.OneBotMessage, args []string) string {
		if len(args) == 0 {
			return "请输入房间号 live del bilibili <房间号>"
		}
		del := c.biliDelLiveService.DelLive(ctx, msg.SelfId, msg.GroupId, args[0])
		return del
	}, model.ScopeGroup)

	// 注册逻辑
	// 子命令 添加逻辑
	liveCmd.AddChild(addCmd)
	addCmd.AddChild(biliCmd)

	// 子命令 删除逻辑
	liveCmd.AddChild(delCmd)
	delCmd.AddChild(biliDelCmd)

	// 注册根命令
	registry.Register(liveCmd)
}
