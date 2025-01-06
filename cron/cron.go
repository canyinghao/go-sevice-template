package cron

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// CronJobManager 用于管理 cron 实例
type CronJobManager struct {
	cron *cron.Cron
}

// NewCronJobManager 创建并返回一个新的 CronJobManager
func NewCronJobManager() *CronJobManager {
	return &CronJobManager{
		cron: cron.New(cron.WithSeconds()),
	}
}

// Start 开始执行定时任务
func (m *CronJobManager) Start() {
	addCron(m)
	m.cron.Start()
}

// Stop 停止执行定时任务
func (m *CronJobManager) Stop() {

	m.cron.Stop()
}

// AddFunc 添加一个新的定时任务
func (m *CronJobManager) AddFunc(schedule string, cmd func()) (cron.EntryID, error) {
	return m.cron.AddFunc(schedule, cmd)
}

// 默认的定时任务配置
func addCron(m *CronJobManager) {

	// m.cron.AddFunc("@every 1h", func() {
	// 	fmt.Println("Task completed at 1s", time.Now())
	// })
	// 支持超时的配置
	m.cron.AddFunc("*/5 * * * * *", func() {

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		go func() {
			time.Sleep(1 * time.Second)
			cancel()
		}()
		defer cancel()
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if err == context.DeadlineExceeded {
				fmt.Println("Task timed out")
			} else {
				fmt.Println("Task finish")
			}
		}
	})

}
