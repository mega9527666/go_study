package socket_connection

import (
	"sync/atomic"

	"github.com/gorilla/websocket"
)

var globalConnID int64 = 0

type Socket_Connection struct {
	Id   int64
	conn *websocket.Conn
}

// 构造函数
func NewSocketConnection(conn *websocket.Conn) *Socket_Connection {
	//每来一个连接，ID 自动 +1，绝对不重复
	// 为什么不用普通的 id++？（重点）
	// 	WebSocket 是并发的：
	// 多个连接
	// 多个 goroutine
	// 会同时执行构造函数
	// 👉 结果：
	// ID 重复
	// 数据错乱
	id := atomic.AddInt64(&globalConnID, 1)
	return &Socket_Connection{
		id:   id,
		conn: conn,
	}
}
