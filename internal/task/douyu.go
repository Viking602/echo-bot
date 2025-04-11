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
		Name:     "checkDouyuLive",
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

	dy := douyu.NewClient()

	all, err := t.douyuSub.GetAll(ctx)
	if err != nil {
		t.log.Errorf("获取所有订阅失败: %v", err)
		return err
	}

	for _, v := range all {
		search, err := dy.Search(strconv.FormatInt(v.RoomId, 10))
		if err != nil {
			t.log.Errorf("获取斗鱼直播间失败: %v", err)
		}

		sub, err := t.sub.GetSubBySubType(ctx, 2)
		if v.LiveState == 0 {

			if search.Data.RecList[0].RoomInfo.IsLive == 1 {

				for _, s := range sub {
					socket, exits := t.bot.GetConnByBotId(s.BotId)
					if !exits {
						t.log.Warn().Err(err).Int64("BotId", s.BotId).Msg("获取 socket 失败")
					}

					t.bot.SendGroupMessage(s.GroupId,
						"直播通知小助手\n"+
							search.Data.RecList[0].RoomInfo.NickName+" 开播啦！\n"+
							"标题："+search.Data.RecList[0].RoomInfo.Description+"\n"+
							"分区："+search.Data.RecList[0].RoomInfo.CateName+"\n"+
							"时间："+time.Unix(search.Data.RecList[0].RoomInfo.LastShowTime, 10).Format("2006-01-02 15:04:05")+"\n"+
							"链接："+search.Data.RecList[0].RoomInfo.BkUrl+"\n"+
							"[CQ:image,file="+search.Data.RecList[0].RoomInfo.RoomSrc+",type=show,id=40000]", socket,
					)
				}

				if err := t.douyuSub.UpdateLive(ctx, &biz.SubDouyuLive{
					RoomId:        v.RoomId,
					LiveState:     1,
					LiveStartTime: search.Data.RecList[0].RoomInfo.LastShowTime,
					LiveEndTime:   0,
				}); err != nil {
					t.log.Error().Err(err).Msg("更新直播间状态失败")
				}
			}
		} else if v.LiveState == 1 {

			endTime := time.Now().Unix()

			if search.Data.RecList[0].RoomInfo.IsLive == 0 {
				for _, s := range sub {
					socket, exits := t.bot.GetConnByBotId(s.BotId)
					if !exits {
						t.log.Warn().Err(err).Int64("BotId", s.BotId).Msg("获取 socket 失败")
					}

					t.bot.SendGroupMessage(s.GroupId,
						"主播："+search.Data.RecList[0].RoomInfo.NickName+" 已下播\n"+
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
