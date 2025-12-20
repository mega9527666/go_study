package hallserver

import (
	"log"
	"mega/common/config"
	"mega/common/db_config"
	"mega/engine/logger"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 生产环境记得校验 origin
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP -> WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("升级失败:", err)
		return
	}
	defer conn.Close()

	logger.Log("客户端连接成功")

	for {
		// 读取消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息失败:", err)
			break
		}

		log.Printf("收到消息: %s\n", msg)

		// 原样返回（echo）
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("发送消息失败:", err)
			break
		}
	}
}

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		logger.Error("初始化端口失败:", os.Args, err)
		return
	}
	env, err := strconv.Atoi(os.Args[2])
	if err != nil {
		logger.Error("初始化环境失败:", os.Args, err)
		return
	}

	config.Environment = env
	db_config.InitDb(config.Environment)
	config.ServerType = config.ServerType_List.Hall_server
	logger.Info("Hall_server.main", os.Args, env, port)

	http.HandleFunc("/ws", wsHandler)
	// log.Println("WebSocket 服务启动: ws://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)

}
