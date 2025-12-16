package account_model

import (
	"mega/common/db_config"
	"mega/engine/logger"
	"mega/engine/mysql_client"
	"mega/engine/mysql_manager"
	"time"
)

// Account 对应 t_accounts 表的结构体
type Account struct {
	ID            int64  `db:"id" json:"id"` // 自增主键
	Account       string `db:"account" json:"account"`
	Pass          string `db:"pass" json:"password"`
	Token         string `db:"token" json:"token"`
	AccountType   int    `db:"account_type" json:"account_type"`
	Status        int    `db:"status" json:"status"`
	IP            string `db:"ip" json:"ip"`
	NickName      string `db:"nick_name" json:"nick_name"`
	Channel       int    `db:"channel" json:"channel"`
	OS            string `db:"os" json:"os"`
	PhoneType     string `db:"phone_type" json:"phone_type"`
	BundleName    string `db:"bundle_name" json:"bundle_name"`
	SystemVersion string `db:"system_version" json:"system_version"`
	CreateTime    int64  `db:"create_time" json:"create_time"`
	LastLoginTime int64  `db:"last_login_time" json:"last_login_time"`
	Phone         string `db:"phone" json:"phone"`
	Sex           int    `db:"sex" json:"sex"`
	Headimgurl    string `db:"headimgurl" json:"headimgurl"`
}

//通道写法
// func IsAccountExist(client mysql_client.Db_client, account string) <-chan struct {
// 	bool
// 	error
// } {
// 	// 创建结果通道
// 	resultChan := make(chan struct {
// 		bool
// 		error
// 	}, 1)

// 	// 启动 goroutine 执行查询
// 	go func() {
// 		var count int
// 		err := client.Db.QueryRow("SELECT COUNT(*) FROM t_accounts WHERE account = ?", account).Scan(&count)

// 		// 发送结果到通道
// 		resultChan <- struct {
// 			bool
// 			error
// 		}{count > 0, err}

// 		close(resultChan)
// 	}()

// 	// 立即返回通道，不阻塞
// 	return resultChan
// }

func IsAccountExist(account string) (bool, error) {
	var client *mysql_client.Db_client = mysql_manager.GetDb(db_config.Db_account, db_config.NowDbType)
	var count int
	err := client.Db.QueryRow("SELECT COUNT(*) FROM t_accounts WHERE account = ?", account).Scan(&count)
	return count > 0, err
}

// Insert 插入账户记录
func InsertAccount(account *Account) (int64, error) {
	var client *mysql_client.Db_client = mysql_manager.GetDb(db_config.Db_account, db_config.NowDbType)
	// 准备 SQL 语句
	sqlStr := `
	INSERT INTO t_accounts (
		account, pass, token, account_type, status, ip, nick_name,
		channel, os, phone_type, bundle_name, system_version,
		create_time, last_login_time, phone, sex, headimgurl
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	// 设置时间戳（如果未设置）
	if account.CreateTime == 0 {
		account.CreateTime = time.Now().UnixMilli()
	}
	if account.LastLoginTime == 0 {
		account.LastLoginTime = time.Now().UnixMilli()
	}

	// 执行插入
	result, err := client.Db.Exec(sqlStr,
		account.Account,
		account.Pass,
		account.Token,
		account.AccountType,
		account.Status,
		account.IP,
		account.NickName,
		account.Channel,
		account.OS,
		account.PhoneType,
		account.BundleName,
		account.SystemVersion,
		account.CreateTime,
		account.LastLoginTime,
		account.Phone,
		account.Sex,
		account.Headimgurl,
	)

	if err != nil {
		logger.Warn("插入账户失败=", err)
	}

	// 获取自增 ID
	id, err := result.LastInsertId()
	if err != nil {
		logger.Warn("获取自增ID失败 =", err)
		return 0, err
	}

	return id, nil
}

func GetAccountByAccount(account string) (*Account, error) {
	var client *mysql_client.Db_client = mysql_manager.GetDb(db_config.Db_account, db_config.NowDbType)
	query := `SELECT id, account, pass,token, account_type, status, ip,nick_name,channel,os,phone_type,bundle_name,system_version,create_time,last_login_time,phone,sex,headimgurl 
              FROM t_accounts WHERE account = ?`

	var acc Account
	err := client.Db.QueryRow(query, account).Scan(
		&acc.ID,
		&acc.Account,
		&acc.Pass,
		&acc.Token,
		&acc.AccountType,
		&acc.Status,
		&acc.IP,
		&acc.NickName,
		&acc.Channel,
		&acc.OS,
		&acc.PhoneType,
		&acc.BundleName,
		&acc.SystemVersion,
		&acc.CreateTime,
		&acc.LastLoginTime,
		&acc.Phone,
		&acc.Sex,
		&acc.Headimgurl,
	)

	if err != nil {
		return nil, err
	}

	return &acc, nil
}
