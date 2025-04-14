package service

import (
	"context"
	"echo/internal/biz"
	"echo/pkg/douyu"
	"echo/pkg/logger"
	"fmt"
	"strconv"
)

type DouyuLiveAddService struct {
	douyu *biz.SubDouyuLiveUsecase
	sub   *biz.SubUsecase
	log   *logger.Logger
}

func NewDouyuLiveAddService(douyu *biz.SubDouyuLiveUsecase, sub *biz.SubUsecase, log *logger.Logger) *DouyuLiveAddService {
	return &DouyuLiveAddService{
		douyu: douyu,
		sub:   sub,
		log:   log,
	}
}

func (s *DouyuLiveAddService) AddLive(ctx context.Context, selfId int64, groupId int64, roomId string) string {
	dy := douyu.NewClient(s.log)
	search, err := dy.GetRoomInfo(roomId)
	if err != nil {
		return "订阅失败"
	}

	room, err := strconv.ParseInt(roomId, 10, 64)
	if err != nil {
		s.log.Errorf("roomId 转换失败: %v", err)
		return "订阅失败"
	}

	if search.Error != 0 {
		return "订阅失败，请检查直播间是否存在"
	}

	subId, err := s.douyu.CreateNewLive(ctx, &biz.SubDouyuLive{
		RoomId: room,
	})
	if err != nil {
		s.log.Errorf("创建订阅失败: %v", err)
		return "订阅失败"
	}

	err, exits := s.sub.CreateSub(ctx, &biz.Sub{
		SubType: 2,
		GroupId: groupId,
		SubId:   subId,
		BotId:   selfId,
		Status:  1,
	})
	if err != nil {
		s.log.Error().Err(err).Int64("selfId", selfId).Int64("groupId", groupId).Msg("创建订阅失败")
		return "订阅失败"
	}

	if exits {
		return fmt.Sprintf("%s 当前订阅记录已存在", roomId)
	}

	return fmt.Sprintf("%s 直播间订阅成功", search.Data.OwnerName)

}
