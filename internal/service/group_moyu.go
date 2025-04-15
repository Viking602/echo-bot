package service

import (
	"context"
	"echo/internal/biz"
	"echo/pkg/logger"
)

type GroupMoyuService struct {
	sub *biz.SubUsecase
	log *logger.Logger
}

func NewGroupMoyuService(sub *biz.SubUsecase, log *logger.Logger) *GroupMoyuService {
	return &GroupMoyuService{
		sub: sub,
		log: log,
	}
}

func (s *GroupMoyuService) GroupMoyu(ctx context.Context, selfId int64, groupId int64, status string) string {
	if status == "on" {
		err, exist := s.sub.CreateSub(ctx, &biz.Sub{
			BotId:   selfId,
			GroupId: groupId,
			SubId:   0,
			SubType: 3,
		})

		if exist {
			return "该群已经开启了此功能"
		}

		if err != nil {
			return "无法开启功能，请联系管理员"
		}
		return "开启功能:摸鱼人日历 成功"
	} else if status == "off" {
		if err := s.sub.DelSub(
			ctx, &biz.Sub{
				BotId:   selfId,
				GroupId: groupId,
				SubId:   0,
				SubType: 3,
			}); err != nil {
			return "无法关闭功能，请联系管理员"
		}
		return "关闭功能:摸鱼人日历 成功"
	}
	return "参数错误"
}
