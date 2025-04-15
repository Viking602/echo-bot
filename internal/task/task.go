package task

import (
	"context"
	"echo/pkg/logger"
	"echo/pkg/scheduler"
	"github.com/google/wire"
	"time"
)

var ProviderSet = wire.NewSet(
	NewTask,
	NewBiliTask,
	NewDouyuTask,
	NewMoyuTask,
)

type RegistrarTask interface {
	Register(scheduler *scheduler.Scheduler) []Info
}

type Info struct {
	Name     string
	Interval time.Duration
	Cron     string
}

// Task 定时任务服务
type Task struct {
	scheduler *scheduler.Scheduler
	log       *logger.Logger
	tasks     []RegistrarTask // 所有任务注册器
	Fn        func(ctx context.Context) error
}

// NewTask 创建定时任务服务
func NewTask(
	log *logger.Logger,
	biliTask *BiliTask,
	douyuTask *DouyuTask,
	moyuTask *MoyuTask,
) *Task {
	sched := scheduler.NewScheduler(log)

	service := &Task{
		scheduler: sched,
		log:       log,
		tasks:     []RegistrarTask{biliTask, douyuTask, moyuTask}, // 添加任务注册器
	}

	// 注册所有定时任务
	service.registerTasks()

	return service
}

// Start 启动所有定时任务
func (s *Task) Start() {
	s.scheduler.Start()
	s.log.Info().Msg("定时任务服务已启动")
}

// Stop 停止所有定时任务
func (s *Task) Stop() {
	s.scheduler.Stop()
	s.log.Info().Msg("定时任务服务已停止")
}

// registerTasks 注册所有定时任务
func (s *Task) registerTasks() {
	var allTaskInfo []Info

	// 收集所有任务信息
	for _, task := range s.tasks {
		// 确保 Register 方法定义在 RegistrarTask 接口中也返回 []InfoTask
		taskInfos := task.Register(s.scheduler)
		allTaskInfo = append(allTaskInfo, taskInfos...)
	}
}
