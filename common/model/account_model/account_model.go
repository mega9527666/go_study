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

// ✅ 纯数据库查询，不需要锁
func IsAccountExist(client mysql_client.Db_client, account string) (bool, error) {
	// 数据库连接是并发安全的（单个连接不支持并发，但连接池支持）
	// sql.DB 内部会处理并发请求
	var count int
	err := client.Db.QueryRow("SELECT COUNT(*) FROM t_accounts WHERE account = ?", account).Scan(&count)
	return count > 0, err
}
