package socket_common

import (
	"mega/engine/http_common"
	"mega/engine/logger"
	"mega/engine/socket_conn_mgr"
	"mega/engine/socket_connection"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 生产环境记得校验 origin
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP -> WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	var ip string = http_common.GetClientIP(r)
	if err != nil {
		logger.Warn("升级失败:", err, ip)
		return
	}
	var socketConn *socket_connection.Socket_Connection = socket_connection.NewSocketConnection(conn, ip)
	socket_conn_mgr.SocketConnManager.AddSocketConnection(socketConn)
	logger.Log("客户端连接成功", ip)
	go socketConn.ReadMsg()
	go socketConn.WritePump()
}
