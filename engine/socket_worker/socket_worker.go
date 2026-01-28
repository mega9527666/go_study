package socket_worker

import (
	"hash/crc32"
	"mega/engine/logger"
	"mega/engine/socket_connection"
)

type MsgTask struct {
	Conn    *socket_connection.Socket_Connection
	MsgType int
	Data    []byte
}

type WorkerPool struct {
	workers []chan MsgTask
}

var GlobalWorkerPool *WorkerPool

func InitWorkerPool(workerNum int) {
	GlobalWorkerPool = NewWorkerPool(workerNum)
}

func NewWorkerPool(workerNum int) *WorkerPool {
	wp := &WorkerPool{
		workers: make([]chan MsgTask, workerNum),
	}

	for i := 0; i < workerNum; i++ {
		ch := make(chan MsgTask, 1024) // 1. 给这个 worker 一个任务队列
		wp.workers[i] = ch

		// 2. 启动一个后台工人
		go func(id int, taskCh chan MsgTask) {
			logger.Log("worker 启动:", id) // 3. 这个工人一辈子只干一件事
			// 4. 不断从自己的队列拿任务
			for task := range taskCh {
				// 5. 顺序处理任务
				task.Conn.OnMessage(task.MsgType, task.Data)
			}
		}(i, ch)
	}

	return wp
}

// ⭐ 按socket id / 连接固定 worker
func (wp *WorkerPool) Dispatch(task MsgTask) {
	idx := int(crc32.ChecksumIEEE(
		[]byte(string(task.Conn.Id)),
	)) % len(wp.workers)

	select {
	case wp.workers[idx] <- task:
	default:
		logger.Warn("worker 队列满，丢弃消息:", task.Conn.Id, task.Data)
	}
}
