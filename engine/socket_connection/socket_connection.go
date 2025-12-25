package socket_connection

import (
	"errors"
	"mega/engine/logger"
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
	Ip     string

	send             chan []byte   // ⭐ 写队列
	closed           chan struct{} // ⭐ 关闭信号
	onMessageHandler MsgHandler
}

type MsgHandler func(s *Socket_Connection, msgType int, data []byte)

// 构造函数
func NewSocketConnection(conn *websocket.Conn, ip string, onMessageHandler MsgHandler) *Socket_Connection {
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
		Id:     id,
		conn:   conn,
		Ip:     ip,
		status: int32(ConnStatusOpen),
		//最多能放 32 个 []byte 元素，不是 32 字节，不是32KB，不是32MB， 32条消息
		send:   make(chan []byte, 32), // 缓冲可调，一般缓存32个字节数组就够了，不会同时发32条消息以上给客户端吧
		closed: make(chan struct{}),   // ⭐ 必须 make
	}
}

func (s *Socket_Connection) ReadMsg() {
	defer s.Close() // ⭐关键
	for {
		// 读取消息
		msgType, msg, err := s.conn.ReadMessage()
		if err != nil {
			logger.Warn("读取消息失败:", s.Ip, err)
			// break
			return
		}

		switch msgType {
		case websocket.TextMessage:
			str := string(msg)
			logger.Log("收到消息 string: ", s.Id, s.Ip, msgType, str)
			// b := []byte(str)
			// s.Send(b)
			s.onMessageHandler(s, msgType, msg)

		case websocket.BinaryMessage:
			logger.Log("收到消息 BinaryMessage: ", s.Id, s.Ip, msgType, msg)
		case websocket.CloseMessage: //websocket.CloseMessage 基本收不到（重要⚠️）大多数情况下： [warn] 读取消息失败: 127.0.0.1 websocket: close 1001 (going away)
			logger.Log("收到消息 CloseMessage: ", s.Id, s.Ip, msgType)
			return
		default:
			logger.Log("收到未知消息类型: ", s.Id, s.Ip, msgType)
		}
	}
}

// 启动写协程（Write Pump）写只能在同一个线程不然有问题的
func (s *Socket_Connection) WritePump() {
	defer s.Close()

	// for msg := range s.send {
	// 	err := s.conn.WriteMessage(websocket.TextMessage, msg)
	// 	if err != nil {
	// 		logger.Warn("写消息失败:", s.Ip, err)
	// 		return
	// 	}
	// }
	for {
		msg, ok := <-s.send
		if !ok {
			// channel 被关闭了
			return
		} else {
			err := s.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				logger.Warn("写消息失败:", s.Ip, err)
				return
			}
		}
	}
}

func (s *Socket_Connection) Send(msg []byte) (err error) {
	if !s.IsOpen() {
		return errors.New("connection closed")
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("connection closed")
		}
	}()
	select {
	case s.send <- msg: //如果缓冲不够就会阻塞
		return nil
	case <-s.closed: //从一个已关闭的 channel 读取，会立刻返回
		return errors.New("connection closed")
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

	close(s.closed) // ⭐ 广播：我关了,不在接受发送给客户端的消息
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
		// ⭐ 关闭 send 通道,不在send消息给客户端了
		close(s.send)
		_ = s.conn.Close()
		s.MarkClosed()
	})
}
