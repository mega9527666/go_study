package mysql_client

import (
	"database/sql"
	dbconfig "mega/common/db_config"
	"mega/engine/logger"
	"time"
)

var DB *sql.DB

func InitDB(dbName string, dbConfig dbconfig.DbConfig) error {
	dsn := dbconfig.GetDBDns(dbConfig, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Warn("InitDB error ", dbName, dbConfig, err)
		return err
	}

	// 连接池配置（非常重要）
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	// 验证连接
	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}
