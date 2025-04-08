package main

import (
	"echo/pkg/logger"
	"github.com/rs/zerolog"
)

func main() {
	logger.SetGlobalLogLevel(zerolog.DebugLevel)

	log := logger.NewLogger()
	app, err := wireApp(log)
	if err != nil {
		panic(err)
	}
	app.Run()
}
