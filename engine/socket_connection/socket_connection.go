package socket_connection

import (
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

// 连接状态
type ConnStatus int32

const (
	ConnStatusOpen ConnStatus = iota
	ConnStatusClosing
	ConnStatusClosed
)

var globalConnID int64 = 0

type Socket_Connection struct {
	Id     int64
	conn   *websocket.Conn
	status int32 // 使用 atomic 操作
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
		Id:   id,
		conn: conn,
	}
}

// 获取状态
func (s *Socket_Connection) Status() ConnStatus {
	//并发安全地读取状态值
	return ConnStatus(atomic.LoadInt32(&s.status))
}

// 是否打开
func (s *Socket_Connection) IsOpen() bool {
	return s.Status() == ConnStatusOpen
}

// 尝试进入 Closing（只允许一次） 如果多线程吧conn关闭就可以根据返回true才是正确关闭
func (s *Socket_Connection) TryClosing() bool {
	//// 状态切换
	return atomic.CompareAndSwapInt32(
		&s.status,
		int32(ConnStatusOpen),
		int32(ConnStatusClosing),
	)
}

// 标记为close
func (s *Socket_Connection) MarkClosed() {
	// 写
	atomic.StoreInt32(&s.status, int32(ConnStatusClosed))
}

func (s *Socket_Connection) Close() {
	// 只允许一个 goroutine 真正关闭
	if !s.TryClosing() {
		return
	}

	// 发送 close frame（可选）
	_ = s.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)

	// _ = s.conn.Close()
	// s.MarkClosed()
	// 给客户端一点时间响应
	//500ms 之后，在一个新的 goroutine 里执行这个函数。
	time.AfterFunc(500*time.Millisecond, func() {
		_ = s.conn.Close()
		s.MarkClosed()
	})
}
