package app

import (
	"github.com/lxzan/gws"
)

type Bot struct {
	handler *Handler
}

func NewBot(handler *Handler) *Bot {
	return &Bot{
		handler: handler,
	}
}

func (b *Bot) Run() error {
	server := gws.NewServer(b.handler, &gws.ServerOption{
		ParallelEnabled: true,
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	})
	if err := server.Run(":6666"); err != nil {
		return err
	}
	return nil
}
