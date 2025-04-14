package task

import (
	"context"
	"echo/internal/biz"
	"echo/internal/bot"
	"echo/pkg/douyu"
	"echo/pkg/logger"
	"echo/pkg/scheduler"
	"echo/utils"
	"strconv"
	"time"
)

type DouyuTask struct {
	log      *logger.Logger
	bot      *bot.Handler
	sub      *biz.SubUsecase
	douyuSub *biz.SubDouyuLiveUsecase
}

func NewDouyuTask(log *logger.Logger, bot *bot.Handler, sub *biz.SubUsecase, douyuSub *biz.SubDouyuLiveUsecase) *DouyuTask {
	return &DouyuTask{
		log:      log,
		bot:      bot,
		sub:      sub,
		douyuSub: douyuSub,
	}
}

func (t *DouyuTask) Register(s *scheduler.Scheduler) []Info {
	task := &scheduler.Task{
		Interval: 30 * time.Second,
		Name:     "斗鱼直播任务",
		Fn:       t.checkDouyuLive,
	}

	s.AddTask(task)

	return []Info{
		{
			Name:     task.Name,
			Interval: task.Interval,
		},
	}
}

func (t *DouyuTask) checkDouyuLive(ctx context.Context) error {

	dy := douyu.NewClient(t.log)

	all, err := t.douyuSub.GetAll(ctx)
	if err != nil {
		t.log.Errorf("获取所有订阅失败: %v", err)
		return err
	}

	for _, v := range all {
		search, err := dy.GetRoomInfo(strconv.FormatInt(v.RoomId, 10))
		if err != nil {
			t.log.Errorf("获取斗鱼直播间失败: %v", err)
		}

		t.log.Info().Int64("error_code", search.Error).Int64("roomId", v.RoomId).Msg("任务检查")

		sub, err := t.sub.GetSubBySubType(ctx, 2)
		if err != nil {
			t.log.Errorf("获取订阅失败: %v", err)
		}

		if search.Error != 0 {
			t.log.Error().Err(err).Msg("斗鱼搜索失败")
			continue
		}

		if v.LiveState == 0 {

			if search.Data.RoomStatus == "1" {

				for _, s := range sub {
					socket, exits := t.bot.GetConnByBotId(s.BotId)
					if !exits {
						t.log.Warn().Err(err).Int64("BotId", s.BotId).Msg("获取 socket 失败")
					}

					t.bot.SendGroupMessage(s.GroupId,
						"直播通知小助手\n"+
							search.Data.OwnerName+" 开播啦！\n"+
							"标题："+search.Data.RoomName+"\n"+
							"分区："+search.Data.CateName+"\n"+
							"时间："+search.Data.StartTime+"\n"+
							"链接："+"https://www.douyu.com/"+search.Data.RoomId+"\n"+
							"[CQ:image,file="+search.Data.RoomThumb+",type=show,id=40000]", socket,
					)
				}

				if err := t.douyuSub.UpdateLive(ctx, &biz.SubDouyuLive{
					RoomId:        v.RoomId,
					LiveState:     1,
					LiveStartTime: utils.ParseTimeToTimestamp(search.Data.StartTime),
					LiveEndTime:   0,
				}); err != nil {
					t.log.Error().Err(err).Msg("更新直播间状态失败")
				}
			}
		} else if v.LiveState == 1 {

			endTime := time.Now().Unix()

			if search.Data.RoomStatus == "2" {
				for _, s := range sub {
					socket, exits := t.bot.GetConnByBotId(s.BotId)
					if !exits {
						t.log.Warn().Err(err).Int64("BotId", s.BotId).Msg("获取 socket 失败")
					}

					t.bot.SendGroupMessage(s.GroupId,
						"主播："+search.Data.OwnerName+" 已下播\n"+
							"直播时长："+utils.FormatDuration(v.LiveStartTime, endTime)+"", socket)
				}

				if err := t.douyuSub.UpdateLive(ctx, &biz.SubDouyuLive{
					RoomId:        v.RoomId,
					LiveState:     0,
					LiveStartTime: 0,
					LiveEndTime:   endTime,
				}); err != nil {
					t.log.Error().Err(err).Msg("更新直播间状态失败")
				}

			}

		}

	}

	return nil

}
