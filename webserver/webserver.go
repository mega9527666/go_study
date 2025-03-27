package main

import (
	"fmt"
	"mega/engine/logger"
	"mega/webreqhandler"
	"os"
	"strconv"
	// "github.com/gorilla/websocket"
)

func main() {
	logger.Log("webserver.main", os.Args[1], fmt.Sprintf("%T", os.Args[1]))
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("初始化端口失败:", os.Args, err)
		return
	}
	logger.Log("webserver.main", os.Args[1], fmt.Sprintf("%T", os.Args[1]), port)
	webreqhandler.ListenAndServe(port)

}
