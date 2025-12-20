package socket_common

import (
	"log"
	"mega/engine/logger"
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
	if err != nil {
		log.Println("升级失败:", err)
		return
	}
	// defer conn.Close()

	logger.Log("客户端连接成功")

	for {
		// 读取消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			logger.Log("读取消息失败:", err)
			break
		}

		logger.Log("收到消息: ", msgType, msg)

		// // 原样返回（echo）
		// err = conn.WriteMessage(msgType, msg)
		// if err != nil {
		// 	logger.Log("发送消息失败:", err)
		// 	break
		// }
	}
}
