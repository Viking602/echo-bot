package service

import (
	"context"
	"echo/internal/biz"
	"fmt"
)

type SetMasterService struct {
	bot *biz.BotUsecase
}

func NewSetMasterService(bot *biz.BotUsecase) *SetMasterService {
	return &SetMasterService{
		bot: bot,
	}
}

func (s *SetMasterService) SetMaster(token string, botId int64, masterId int64) string {
	ctx := context.Background()

	botInfo, err := s.bot.GetBot(ctx, botId)
	if err != nil {
		return "设置失败"
	}
	if token == botInfo.AccessToken && botInfo.SelfId == 0 {
		if err := s.bot.UpdateSelfId(context.Background(), masterId, botId); err != nil {
			return "设置失败"
		}
		return "设置成功"
	}

	if botInfo.SelfId != 0 {
		return fmt.Sprintf("已设置过主人:%v", botInfo.SelfId)
	}

	return "设置失败"
}
