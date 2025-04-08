package biz

import (
	"context"
	"fmt"
)

type Bot struct {
	Id         int64
	BotId      int64
	BotName    string
	SelfId     int64
	Status     int
	CreateTime string
	UpdateTime string
}

type BotRepo interface {
	Create(ctx context.Context, bot *Bot) error
	GetBot(ctx context.Context, botId int64) (*Bot, error)
}

type BotUsecase struct {
	repo BotRepo
}

func NewBotUsecase(botRepo BotRepo) *BotUsecase {
	return &BotUsecase{
		repo: botRepo,
	}
}

func (u *BotUsecase) CreateBot(ctx context.Context, botId int64) error {
	// 检查是否已经存在
	_, err := u.repo.GetBot(ctx, botId)
	if err == nil {
		return err
	}

	if err := u.repo.Create(ctx, &Bot{
		BotId:   botId,
		BotName: fmt.Sprintf("bot_%d", botId),
	}); err != nil {
		return err
	}

	return nil
}
