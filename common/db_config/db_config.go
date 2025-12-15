package dbconfig

import (
	"mega/engine/logger"
)

type dbType int

const (
	DbType_Dev    dbType = 1
	DbType_Test   dbType = 2
	DbType_Online dbType = 3
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

var NowDbType dbType = DbType_Dev

// ==================
// DB 配置表
// ==================

var dbMap = map[dbType]*DbConfig{
	DbType_Dev: {
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pass: "f301517757c8cf6c",
	},
	DbType_Test: {
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pass: "f301517757c8cf6c",
	},
	DbType_Online: {
		Host: "127.0.0.1",
		Port: 3306,
		User: "root",
		Pass: "f301517757c8cf6c",
	},
}

// ==================
// 初始化
// ==================

func InitDb(dbType dbType) {
	NowDbType = dbType
}

// ==================
// 获取 DB 配置
// ==================
func GetDbConfig(dbType dbType) *DbConfig {
	db, ok := dbMap[dbType]
	if !ok {
		logger.Warn("GetDbConfig error 不存在数据库类型", dbType, dbMap)
	}
	return db
}
