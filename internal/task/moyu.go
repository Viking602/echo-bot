package task

import (
	"context"
	"echo/internal/biz"
	"echo/internal/bot"
	"echo/pkg/logger"
	"echo/pkg/moyu"
	"echo/pkg/scheduler"
)

type MoyuTask struct {
	log *logger.Logger
	bot *bot.Handler
	sub *biz.SubUsecase
}

func NewMoyuTask(log *logger.Logger, bot *bot.Handler, sub *biz.SubUsecase) *MoyuTask {
	return &MoyuTask{
		log: log,
		bot: bot,
		sub: sub,
	}
}

func (t *MoyuTask) Register(s *scheduler.Scheduler) []Info {

	task := &scheduler.Task{
		Name: "摸鱼人日历",
		Cron: "0 0 8 * * *",
		Fn:   t.checkMoyu,
	}

	s.AddTask(task)

	return []Info{
		{
			Cron: task.Cron,
			Name: "摸鱼人日历",
		},
	}
}

func (t *MoyuTask) checkMoyu(ctx context.Context) error {

	moyuClient := moyu.NewClient()

	sub, err := t.sub.GetSubBySubIdSubType(ctx, 0, 3)
	if err != nil {
		t.log.Error().Err(err).Msg("获取订阅失败")
		return err
	}

	for _, s := range sub {
		socket, exits := t.bot.GetConnByBotId(s.BotId)
		if !exits {
			t.log.Warn().Err(err).Int64("BotId", s.BotId).Msg("获取 socket 失败")
			continue
		}

		api, err := moyuClient.FishAPI()
		if err != nil {
			t.log.Warn().Err(err).Msg("获取摸鱼日历失败")
			continue
		}

		if api.Success {
			t.bot.SendGroupMessage(s.GroupId, "[CQ:image,file="+api.Url+",type=show,id=40000]", socket)
		}

	}

	return nil
}
