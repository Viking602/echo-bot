package biz

import (
	"context"
	"fmt"
	"time"
)

type Bot struct {
	Id             int64
	BotId          int64
	BotName        string
	SelfId         int64
	Status         int
	LastOnlineTime time.Time
	LastOnlineIP   string
	CreateTime     string
	UpdateTime     string
}

type BotRepo interface {
	Create(ctx context.Context, bot *Bot) error
	GetBot(ctx context.Context, botId int64) (*Bot, error)
	UpdateUptime(ctx context.Context, bot *Bot) error
}

type BotUsecase struct {
	repo BotRepo
}

func NewBotUsecase(botRepo BotRepo) *BotUsecase {
	return &BotUsecase{
		repo: botRepo,
	}
}

func (u *BotUsecase) CreateBot(ctx context.Context, botId int64, uptime int64, ip string) error {
	// 检查是否已经存在
	_, err := u.repo.GetBot(ctx, botId)
	if err == nil {
		return err
	}

	if err := u.repo.Create(ctx, &Bot{
		BotId:          botId,
		LastOnlineTime: time.Unix(uptime, 0),
		LastOnlineIP:   ip,
		BotName:        fmt.Sprintf("bot_%d", botId),
	}); err != nil {
		return err
	}

	return nil
}

func (u *BotUsecase) UpdateBotUptime(ctx context.Context, botId int64, uptime int64, ip string) error {
	if err := u.repo.UpdateUptime(ctx, &Bot{
		BotId:          botId,
		LastOnlineTime: time.Unix(uptime, 0),
		LastOnlineIP:   ip,
	}); err != nil {
		return err
	}

	return nil
}
