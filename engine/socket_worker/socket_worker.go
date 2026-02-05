package socket_worker

import (
	"mega/engine/logger"
)

type Task func()

type WorkerPool struct {
	workers []chan Task
}

var GlobalWorkerPool *WorkerPool

func InitWorkerPool(workerNum int) {
	GlobalWorkerPool = NewWorkerPool(workerNum)
}

func NewWorkerPool(workerNum int) *WorkerPool {
	wp := &WorkerPool{
		workers: make([]chan Task, workerNum),
	}

	for i := 0; i < workerNum; i++ {
		ch := make(chan Task, 1024)
		wp.workers[i] = ch

		go func(id int, taskCh chan Task) {
			logger.Log("worker 启动:", id)
			for task := range taskCh {
				task()
			}
		}(i, ch)
	}
	return wp
}

// ⭐ key 用来保证同一连接固定落到同一个 worker
func (wp *WorkerPool) Dispatch(id int64, task Task) {
	idx := int(id % int64(len(wp.workers)))

	select {
	case wp.workers[idx] <- task:
	default:
		logger.Warn("worker 队列满，丢弃任务:", id)
	}
}
