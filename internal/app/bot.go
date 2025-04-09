package app

import (
	"github.com/lxzan/gws"
	"net/http"
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

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token != "" {
			if len(token) > 7 && token[:7] == "Bearer " {
				token = token[7:]
			}
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		upgrader := gws.NewUpgrader(b.handler, &gws.ServerOption{
			ParallelEnabled:   true,
			Recovery:          gws.Recovery,
			PermessageDeflate: gws.PermessageDeflate{Enabled: true},
		})
		socket, err := upgrader.Upgrade(w, r)
		if err != nil {
			return
		}

		socket.Session().Store("token", token)

		go socket.ReadLoop()
	})

	if err := http.ListenAndServe(":6666", nil); err != nil {
		return err
	}
	return nil
}
