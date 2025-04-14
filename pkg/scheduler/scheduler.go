package scheduler

import (
	"context"
	"echo/pkg/logger"
	"sync"
	"time"
)

type Task struct {
	Name     string                          // 任务名称
	Interval time.Duration                   // 执行间隔
	Fn       func(ctx context.Context) error // 执行函数
}

// Scheduler 定时任务调度器
type Scheduler struct {
	tasks     map[string]*Task
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
	isRunning bool
	mu        sync.Mutex
	log       *logger.Logger
}

// NewScheduler 创建一个新的定时任务调度器
func NewScheduler(log *logger.Logger) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &Scheduler{
		tasks:  make(map[string]*Task),
		ctx:    ctx,
		cancel: cancel,
		log:    log,
	}
}

// AddTask 添加一个定时任务
func (s *Scheduler) AddTask(task *Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.Name] = task
}

// RemoveTask 移除一个定时任务
func (s *Scheduler) RemoveTask(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[name]; exists {
		delete(s.tasks, name)
		s.log.Info().Str("task", name).Msg("移除定时任务")
	}
}

// Start 启动定时任务调度器
func (s *Scheduler) Start() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRunning {
		return
	}

	s.isRunning = true
	s.log.Info().Msg("定时任务调度器启动")

	for name, task := range s.tasks {
		s.wg.Add(1)
		go s.runTask(name, task)
	}
}

// Stop 停止定时任务调度器
func (s *Scheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		return
	}

	s.cancel()
	s.wg.Wait()
	s.isRunning = false
	s.log.Info().Msg("定时任务调度器停止")
}

// runTask 运行一个定时任务
func (s *Scheduler) runTask(name string, task *Task) {
	defer s.wg.Done()
	defer func() {
		if r := recover(); r != nil {
			s.log.Error().Interface("panic", r).Str("task", name).Msg("定时任务发生异常")
		}
	}()

	s.log.Info().Msgf("定时任务启动: %s, 间隔: %s", name, task.Interval.String())

	for {
		// 检查上下文是否取消
		if s.ctx.Err() != nil {
			s.log.Info().Str("task", name).Msg("定时任务停止")
			return
		}

		// 等待指定时间
		select {
		case <-s.ctx.Done():
			s.log.Info().Str("task", name).Msg("定时任务停止")
			return
		case <-time.After(task.Interval):
			// 继续执行
		}

		// 执行任务
		s.log.Debug().Str("task", name).Str("time", time.Now().Format("15:04:05")).Msg("执行定时任务")
		if err := task.Fn(s.ctx); err != nil {
			s.log.Error().Err(err).Str("task", name).Msg("执行定时任务出错")
		}
	}
}
