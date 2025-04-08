package bot

import (
	"github.com/lxzan/gws"
)

type Bot struct {
	handler *Handler
}

func NewBot(handler *Handler) *Bot {
	// 在 Bot 初始化时注册指令
	return &Bot{
		handler: handler,
	}
}

func (b *Bot) Run() {
	server := gws.NewServer(b.handler, &gws.ServerOption{
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	})
	server.Run(":6666")
}
