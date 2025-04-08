package data

import (
	"context"
	"echo/internal/biz"
	"echo/internal/data/ent"
	bot2 "echo/internal/data/ent/bot"
	"time"
)

type botRepo struct {
	data *Data
}

func NewBotRepo(data *Data) biz.BotRepo {
	return &botRepo{data: data}
}

func (r *botRepo) Create(ctx context.Context, bot *biz.Bot) error {
	_, err := r.data.db.Bot.Create().
		SetBotName(bot.BotName).
		SetBotID(bot.BotId).
		SetSelfID(bot.SelfId).
		SetStatus(1).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *botRepo) GetBot(ctx context.Context, botId int64) (*biz.Bot, error) {
	bot, err := r.data.db.Bot.Query().Where(bot2.BotID(botId)).First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return &biz.Bot{
		Id:         bot.ID,
		BotId:      bot.BotID,
		BotName:    bot.BotName,
		SelfId:     bot.SelfID,
		Status:     bot.Status,
		CreateTime: bot.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: bot.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}
