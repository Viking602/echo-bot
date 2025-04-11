package service

import (
	"context"
	"echo/internal/biz"
	"echo/pkg/bilibili"
	"echo/pkg/logger"
	"fmt"
	"strconv"
)

type BiliLiveAddService struct {
	bili *biz.SubBiliLiveUsecase
	sub  *biz.SubUsecase
	log  *logger.Logger
}

func NewBiliLiveAddService(bili *biz.SubBiliLiveUsecase, sub *biz.SubUsecase, log *logger.Logger) *BiliLiveAddService {
	return &BiliLiveAddService{
		bili: bili,
		sub:  sub,
		log:  log,
	}
}

func (s *BiliLiveAddService) AddLive(ctx context.Context, selfId int64, groupId int64, roomId string) string {
	bili := bilibili.NewClient()

	live, err := bili.GetLiveState(roomId)
	if err != nil {
		s.log.Error().Err(err).Int64("selfId", selfId).Int64("groupId", groupId).Msg("获取直播间信息失败")
		return "获取直播间信息失败"
	}

	if live.Code != 0 {
		return live.Message
	}

	info, err := bili.GetMasterInfo(strconv.FormatInt(live.Data.Uid, 10))
	if err != nil {
		s.log.Error().Err(err).Int64("selfId", selfId).Int64("groupId", groupId).Msg("获取主播信息失败")
		return "获取主播信息失败"
	}

	createId, err := s.bili.CreateNewBiliLive(ctx, &biz.SubBiliLive{
		RoomId:        live.Data.RoomId,
		LiveState:     0,
		LiveStartTime: 0,
		LiveEndTime:   0,
	})
	if err != nil {
		s.log.Error().Err(err).Int64("selfId", selfId).Int64("groupId", groupId).Msg("创建直播间订阅失败")
		return "订阅失败"
	}

	err, exits := s.sub.CreateSub(ctx, &biz.Sub{
		GroupId: groupId,
		SubId:   createId,
		BotId:   selfId,
		SubType: 1,
		Status:  1,
	})
	if err != nil {
		s.log.Error().Err(err).Int64("selfId", selfId).Int64("groupId", groupId).Msg("创建订阅失败")
		return "订阅失败"
	}

	if exits {
		return fmt.Sprintf("%s 当前订阅记录已存在", roomId)
	}

	return fmt.Sprintf("%s 直播间订阅成功", info.Data.Info.Uname)

}
