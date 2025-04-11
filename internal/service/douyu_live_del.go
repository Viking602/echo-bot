package service

import (
	"context"
	"echo/internal/biz"
	"echo/pkg/logger"
	"strconv"
)

type DouyuLiveDelService struct {
	douyu *biz.SubDouyuLiveUsecase
	sub   *biz.SubUsecase
	log   *logger.Logger
}

func NewDouyuLiveDelService(douyu *biz.SubDouyuLiveUsecase, sub *biz.SubUsecase, log *logger.Logger) *DouyuLiveDelService {
	return &DouyuLiveDelService{
		douyu: douyu,
		sub:   sub,
		log:   log,
	}
}

func (s *DouyuLiveDelService) DelLive(ctx context.Context, selfId int64, groupId int64, roomId string) string {
	room, err := strconv.ParseInt(roomId, 10, 64)
	if err != nil {
		s.log.Errorf("roomId 转换失败: %v", err)
		return "取消订阅失败，请联系管理员"
	}

	live, err := s.douyu.GetByRoomId(ctx, room)
	if err != nil {
		s.log.Errorf("获取直播间失败: %v", err)
		return "取消订阅失败，请确认订阅是否存在"
	}

	subAll, err := s.sub.GetAllBySubIdSubType(ctx, live.Id, 1)
	if err != nil {
		s.log.Errorf("获取订阅记录失败: %v", err)
		return "取消订阅失败，请联系管理员"
	}

	if err := s.sub.DelSub(ctx, &biz.Sub{
		BotId:   selfId,
		GroupId: groupId,
		SubId:   live.Id,
		SubType: 2,
	}); err != nil {
		s.log.Errorf("删除订阅失败: %v", err)
		return "取消订阅失败，请确认订阅是否存在"
	}

	if len(subAll) == 1 {
		if err := s.douyu.Del(ctx, live.Id); err != nil {
			s.log.Errorf("删除直播间记录失败: %v", err)
			return "取消订阅失败，请联系管理员"
		}
	}

	return "取消订阅成功"

}
