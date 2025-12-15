package account_model

import "mega/engine/mysql_client"

// Account 对应 t_accounts 表的结构体
type Account struct {
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

// 回调函数方式访问异步sql操作
func IsAccountExist(client *mysql_client.Db_client, account string, callback func(bool, error)) {
	go func() {
		var count int
		err := client.Db.QueryRow("SELECT COUNT(*) FROM t_accounts WHERE account = ?", account).Scan(&count)
		// 调用回调函数
		callback(count > 0, err)
	}()
}
