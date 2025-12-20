package hallserver

import (
	"mega/common/config"
	"mega/common/db_config"
	"mega/engine/logger"
	"os"
	"strconv"
)

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

}
