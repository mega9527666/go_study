package mysql_client

import (
	"database/sql"
	"mega/common/db_config"
	"mega/engine/logger"
	"time"

	// ✅ 必须导入 MySQL 驱动
	_ "github.com/go-sql-driver/mysql"
)

type Db_client struct {
	DbName   string
	DbConfig *db_config.DbConfig
	Db       *sql.DB
}

// 构造函数
func newDbClient(dbName string, config *db_config.DbConfig, db *sql.DB) *Db_client {
	return &Db_client{
		DbName:   dbName,
		DbConfig: config,
		Db:       db,
	}
}

func InitDB(dbName string, dbConfig *db_config.DbConfig) (*Db_client, error) {
	dsn := db_config.GetDBDns(dbConfig, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Warn("InitDB error ", dbName, dbConfig, err)
		return nil, err
	}

	// 连接池配置（非常重要）
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	// 验证连接  sql.Open 不会立刻连数据库，Ping() 才会。
	if err := db.Ping(); err != nil {
		logger.Warn("InitDB ping error ", dbName, dbConfig, err)
		return nil, err
	}

	var client *Db_client = newDbClient(dbName, dbConfig, db)
	return client, nil
}
