//go:build wireinject
// +build wireinject

package main

import (
	"echo/internal/app"
	"echo/internal/biz"
	"echo/internal/command"
	"echo/internal/data"
	"echo/internal/service"
	"echo/pkg/logger"
	"github.com/google/wire"
)

func wireApp(log *logger.Logger) (*app.Bot, error) {
	panic(wire.Build(
		app.NewBotHandler,   // 提供 *bot.BotHandler
		app.NewBot,          // 提供 *bot.Bot
		service.ProviderSet, // 提供 service 包的依赖
		command.ProviderSet, // 提供 command 包的依赖
		data.ProviderSet,    // 提供 data 包的依赖
		biz.ProviderSet,     // 提供 biz 包的依赖
	))
}
