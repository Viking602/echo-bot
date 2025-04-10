//go:build wireinject
// +build wireinject

package main

import (
	"echo/internal/app"
	"echo/internal/biz"
	"echo/internal/bot"
	"echo/internal/command"
	"echo/internal/data"
	"echo/internal/service"
	"echo/internal/task"
	"echo/pkg/logger"
	"github.com/google/wire"
)

func wireApp(log *logger.Logger) (*app.App, error) {
	panic(wire.Build(
		bot.NewBotHandler, // 提供 *bot.BotHandler
		bot.NewBot,        // 提供 *bot.Bot
		app.NewApp,
		service.ProviderSet, // 提供 service 包的依赖
		command.ProviderSet, // 提供 command 包的依赖
		data.ProviderSet,    // 提供 data 包的依赖
		biz.ProviderSet,     // 提供 biz 包的依赖
		task.ProviderSet,    // 提供 task 包的依赖
	))
}
