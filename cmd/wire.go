//go:build wireinject
// +build wireinject

package main

import (
	"echo/internal/command"
	"echo/internal/service"
	"echo/pkg/app"
	bot2 "echo/pkg/bot"
	"echo/pkg/logger"
	"github.com/google/wire"
)

func wireApp(log *logger.Logger) (*app.App, error) {
	panic(wire.Build(
		bot2.NewBotHandler,  // 提供 *bot.BotHandler
		bot2.NewBot,         // 提供 *bot.Bot
		service.ProviderSet, // 提供 service 包的依赖
		command.ProviderSet, // 提供 command 包的依赖
		app.NewApp,          // 提供 *app.App
	))
}
