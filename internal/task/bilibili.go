package task

import (
	"context"
	"echo/internal/biz"
	"echo/internal/bot"
	"echo/pkg/bilibili"
	"echo/pkg/logger"
	"echo/pkg/scheduler"
	"echo/utils"
	"strconv"
	"time"
)

type BiliTask struct {
	log     *logger.Logger
	bot     *bot.Handler
	sub     *biz.SubUsecase
	biliSub *biz.SubBiliLiveUsecase
}

func NewBiliTask(log *logger.Logger, bot *bot.Handler, sub *biz.SubUsecase, biliSub *biz.SubBiliLiveUsecase) *BiliTask {
	return &BiliTask{
		log:     log,
		bot:     bot,
		sub:     sub,
		biliSub: biliSub,
	}
}

func (t *BiliTask) Register(s *scheduler.Scheduler) []Info {
	task := &scheduler.Task{
		Name:     "check-bilibili-live",
		Interval: time.Second * 60,
		Fn:       t.checkBiliLive,
	}

	s.AddTask(task)

	return []Info{
		{
			Name:     task.Name,
			Interval: task.Interval,
		},
	}

}

func (t *BiliTask) checkBiliLive(ctx context.Context) error {

	bili := bilibili.NewClient()

	all, err := t.biliSub.GetAll(ctx)
	if err != nil {
		t.log.Error().Err(err).Msg("获取数据失败")
		return err
	}

	for _, v := range all {
		if v.LiveState == 0 {
			state, err := bili.GetLiveState(strconv.FormatInt(v.RoomId, 10))
			if err != nil {
				t.log.Warn().Err(err).Msg("获取直播间状态失败")
			}

			if state.Code == 0 && state.Data.LiveStatus == 1 {
				sub, err := t.sub.GetSubBySubId(ctx, v.Id)
				if err != nil {
					t.log.Warn().Err(err).Msg("获取数据失败")
				}
				for _, s := range sub {
					socket, exits := t.bot.GetConnByBotId(s.BotId)
					if !exits {
						t.log.Warn().Err(err).Int64("BotId", s.BotId).Msg("获取 socket 失败")
					}

					master, err := bili.GetMasterInfo(strconv.FormatInt(state.Data.Uid, 10))
					if err != nil {
						t.log.Warn().Err(err).Msg("获取主播信息失败")
					}
					t.bot.SendGroupMessage(s.GroupId,
						"直播通知小助手\n"+
							"主播："+master.Data.Info.Uname+" 开始直播啦！\n"+
							"标题："+state.Data.Title+"\n"+
							"分区："+state.Data.AreaName+"\n"+
							"时间："+state.Data.LiveTime+"\n"+
							"链接：https://live.bilibili.com/"+strconv.FormatInt(int64(master.Data.RoomId), 10)+"\n"+
							"[CQ:image,file="+state.Data.UserCover+",type=show,id=40000]", socket,
					)

					if err := t.biliSub.UpdateBiliLive(ctx, &biz.SubBiliLive{
						RoomId:        v.RoomId,
						LiveState:     state.Data.LiveStatus,
						LiveStartTime: utils.ParseTimeToTimestamp(state.Data.LiveTime),
					}); err != nil {
						t.log.Error().Err(err).Msg("更新直播间状态失败")
					}

				}

			}
		} else if v.LiveState == 1 {
			state, err := bili.GetLiveState(strconv.FormatInt(v.RoomId, 10))
			if err != nil {
				t.log.Warn().Err(err).Msg("获取直播间状态失败")
			}
			if state.Code == 0 {
				if err := t.biliSub.UpdateBiliLive(ctx, &biz.SubBiliLive{
					RoomId:      v.RoomId,
					LiveState:   state.Data.LiveStatus,
					LiveEndTime: time.Now().Unix(),
				}); err != nil {
					t.log.Error().Err(err).Msg("更新直播间状态失败")
				}
			}
		}
	}
	return nil
}
