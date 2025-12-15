package mysql_manager

import (
	"mega/common/db_config"
	"mega/common/mysql_client"
	"mega/engine/logger"
	"strconv"
	"sync"
)

// 定义一个全局的 map 变量
var (
	dbMap     = make(map[string]*mysql_client.Db_client)
	dbMapLock sync.RWMutex // // 如果读操作远多于写操作，使用 RWMutex 性能更好 // 如果读写比例差不多，使用 Mutex 更简单
)

func GetCodeDbName(dbName string, dbType int) string {
	var realDbName string = dbName + "_db_" + strconv.Itoa(dbType)
	return realDbName
}

// GetDbClient 根据 key 获取 Db_client 实例
func getDbClient(key string) (*mysql_client.Db_client, bool) {
	dbMapLock.RLock()            //获取读锁（允许多个 goroutine 同时获取）
	defer dbMapLock.RUnlock()    //// 第2步：延迟释放读锁
	client, exists := dbMap[key] //// 第3步：执行读操作（在锁的保护下）
	// 函数返回前的一瞬间，defer 执行！
	return client, exists
}

// SetDbClient 设置或更新 Db_client 实例
func setDbClient(key string, client *mysql_client.Db_client) {
	dbMapLock.Lock()
	defer dbMapLock.Unlock()
	dbMap[key] = client
}

// 获取 DB
func GetDb(dbName string, dbType int) *mysql_client.Db_client {
	if dbType == 0 {
		dbType = db_config.NowDbType
	}

	realDbName := GetCodeDbName(dbName, dbType)
	client, exists := getDbClient(realDbName)
	if exists {
		return client
	} else {
		dbConfig := db_config.GetDbConfig(dbType)
		client, err := mysql_client.InitDB(dbName, dbConfig)
		if err != nil {
			logger.Warn("GetDb error", err)
			return nil
		}
		setDbClient(realDbName, client)
		return client
	}

}
