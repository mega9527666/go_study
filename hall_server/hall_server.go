package hallserver

import (
	"mega/common/config"
	"mega/common/db_config"
	"mega/engine/logger"
	"mega/engine/socket_common"
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

	// http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/ws", socket_common.WsHandler)
	// log.Println("WebSocket 服务启动: ws://localhost:8080/ws")
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
