package app

import (
	"echo/pkg/bot"
	"echo/pkg/logger"
)

type App struct {
	bot    *bot.Bot
	logger *logger.Logger
}

func NewApp(bot *bot.Bot, logger *logger.Logger) *App {

	return &App{
		bot:    bot,
		logger: logger,
	}
}

func (a *App) Run() {
	a.logger.Info().
		Str("app", "echo").
		Msg("Starting bot")
	a.bot.Run()
}
