package app

import (
	"echo/internal/bot"
	"echo/internal/task"
)

type App struct {
	bot  *bot.Bot
	task *task.Task
}

func NewApp(bot *bot.Bot, task *task.Task) *App {
	return &App{
		bot:  bot,
		task: task,
	}
}

func (a *App) Run() error {

	defer a.task.Stop()

	a.task.Start()

	if err := a.bot.Run(); err != nil {
		return err
	}
	return nil
}
