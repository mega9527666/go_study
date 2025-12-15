package dbconfig

import (
	"mega/engine/logger"
)

const (
	DbType_Dev    int = 1
	DbType_Test   int = 2
	DbType_Online int = 3
)

// ==================
// DB Name 常量
// ==================

const (
	Db_account = "db_account"
	Db_game    = "db_game"
)

// ==================
// DbConfig 结构体
// ==================

type DbConfig struct {
	Host string
	Port int
	User string
	Pass string
}

// ==================
// 当前 DB 类型
// ==================

var NowDbType int = DbType_Dev

// ==================
// DB 配置表
// ==================

var dbMap = map[int]*DbConfig{
	DbType_Dev: {
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pass: "Mega@9527",
	},
	DbType_Test: {
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pass: "Mega@9527",
	},
	DbType_Online: {
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pass: "Mega@9527",
	},
}

// ==================
// 初始化
// ==================

func InitDb(dbType int) {
	NowDbType = dbType
}

// ==================
// 获取 DB 配置
// ==================
func GetDbConfig(dbType int) *DbConfig {
	db, ok := dbMap[dbType]
	if !ok {
		logger.Warn("GetDbConfig error 不存在数据库类型", dbType, dbMap)
	}
	return db
}
