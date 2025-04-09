package main

import (
	"echo/pkg/logger"
	"github.com/rs/zerolog"
)

func main() {
	logger.SetGlobalLogLevel(zerolog.InfoLevel)

	log := logger.NewLogger()
	app, err := wireApp(log)
	if err != nil {
		panic(err)
	}

	log.Info().Msg("echo-bot 启动中...")
	if err := app.Run(); err != nil {
		panic(err)
	}
}
