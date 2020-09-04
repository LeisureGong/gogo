package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner在给定的超市时间内执行一组任务
// 并在OS发送中断信号时结束这些任务
type Runner struct {
	// 接收中断信号的通道
	interrupt chan os.Signal

	// 报告处理任务已经完成
	complete chan error

	// 	报告处理任务已经超时
	timeout <-chan time.Time

	// 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// timeout log
var ErrTimeout = errors.New("超时")

// interrupt log
var ErrInterrupt = errors.New("中断")

// 返回一个新Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err

	case <-r.timeout:
		return ErrTimeout
	}
}

// 	执行每一个注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测OS的中断信号
		if r.goInterrupt() {
			return ErrInterrupt
		}

		// 执行一注册的的任务
		task(id)
	}
	return nil
}

// 验证是否接收到了中断信号
func (r *Runner) goInterrupt() bool {
	select {
	// 	当中断时间被触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true

	// 继续正常运行
	default:
		return false
	}
}
