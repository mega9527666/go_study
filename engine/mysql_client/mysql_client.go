package mysql_client

import (
	"database/sql"
	dbconfig "mega/common/db_config"
	"mega/engine/logger"
	"time"
)

type Db_client struct {
	DbName   string
	DbConfig *dbconfig.DbConfig
	Db       *sql.DB
}

// 构造函数
func newDbClient(dbName string, config *dbconfig.DbConfig, db *sql.DB) *Db_client {
	return &Db_client{
		DbName:   dbName,
		DbConfig: config,
		Db:       db,
	}
}

func InitDB(dbName string, dbConfig *dbconfig.DbConfig) (*Db_client, error) {
	dsn := dbconfig.GetDBDns(dbConfig, dbName)

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
