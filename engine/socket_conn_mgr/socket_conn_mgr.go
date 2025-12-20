package socket_conn_mgr

import (
	"mega/engine/logger"
	"mega/engine/socket_connection"
	"sync"
)

type SocketConnectionManager struct {
	socketConnMap map[int64]*socket_connection.Socket_Connection
	mu            sync.RWMutex
}

// 单例
var SocketConnManager = &SocketConnectionManager{
	socketConnMap: make(map[int64]*socket_connection.Socket_Connection),
}

// 增加客户端 socket 连接
func (m *SocketConnectionManager) AddSocketConnection(conn *socket_connection.Socket_Connection) {
	if conn == nil {
		logger.Warn("加个空的conn")
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	m.socketConnMap[conn.Id] = conn
}

// 根据 id 获取连接
func (m *SocketConnectionManager) GetSocketConnection(id int64) *socket_connection.Socket_Connection {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.socketConnMap[id]
}

// 删除客户端 socket 连接
func (m *SocketConnectionManager) RemoveSocketConnection(removeOne *socket_connection.Socket_Connection) {
	if removeOne == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.socketConnMap, removeOne.Id)
}
