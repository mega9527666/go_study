package config

import (
	"errors"
	"mega/engine/logger"
	"mega/engine/random_util"
	"strconv"

	"github.com/spf13/viper"
)

const (
	EnvDev    int = 1
	EnvTest   int = 2
	EnvOnline int = 3
)

var Env = EnvDev

type serverType_struct struct {
	Web_server     string
	Account_server string
	Hall_server    string
}

var ServerType_List = serverType_struct{
	Web_server:     "[web_server]",
	Account_server: "[account_server]",
	Hall_server:    "[hall_server]",
}

var ServerType = ServerType_List.Account_server

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ServerItem struct {
	ID         int    `mapstructure:"id"`
	IP         string `mapstructure:"ip"`
	InternalIP string `mapstructure:"internalIp"`
	HttpPort   int    `mapstructure:"http_port"`
	GrpcPort   int    `mapstructure:"grpc_port"`
	SocketPort int    `mapstructure:"socket_port"`
}

type Config struct {
	DB            DBConfig     `mapstructure:"db"`
	WebServer     []ServerItem `mapstructure:"web_server"`
	AccountServer []ServerItem `mapstructure:"account_server"`
	HallServer    []ServerItem `mapstructure:"hall_server"`
}

var Global_Config Config //全局配置
var Now_ServerItem ServerItem

func envToConfigName(env int) string {
	switch env {
	case EnvDev:
		return "dev"
	case EnvTest:
		return "test"
	case EnvOnline:
		return "online"
	default:
		return "dev"
	}
}

func findServerByHttpPort(list []ServerItem, port int) (ServerItem, bool) {
	for _, item := range list {
		if item.HttpPort == port {
			return item, true
		}
	}
	return ServerItem{}, false
}

func InitNowServerItem(http_port int) error {
	var (
		item ServerItem
		ok   bool
	)

	switch ServerType {
	case ServerType_List.Web_server:
		item, ok = findServerByHttpPort(Global_Config.WebServer, http_port)

	case ServerType_List.Account_server:
		item, ok = findServerByHttpPort(Global_Config.AccountServer, http_port)

	case ServerType_List.Hall_server:
		item, ok = findServerByHttpPort(Global_Config.HallServer, http_port)

	default:
		logger.Error("unknown ServerType", ServerType)
		return errors.New("unknown ServerType")
	}

	if !ok {
		logger.Error("http_port not found in config")
		return errors.New("http_port not found in config")
	}

	Now_ServerItem = item
	logger.Log("InitNowServerItem ID=", Now_ServerItem.ID)
	logger.Log("InitNowServerItem HttpPort=", Now_ServerItem.HttpPort)
	logger.Log("InitNowServerItem GrpcPort=", Now_ServerItem.GrpcPort)
	logger.Log("InitNowServerItem SocketPort=", Now_ServerItem.SocketPort)
	return nil
}

func InitConfig(http_port int) error {
	configName := envToConfigName(Env)

	viper.Reset()
	viper.SetConfigName(configName) // dev / test / online
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	// 支持环境变量
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Warn("InitConfig ReadInConfig error=", err)
		return err
	}

	if err := viper.Unmarshal(&Global_Config); err != nil {
		logger.Warn("InitConfig Unmarshal error=", err)
		return err
	}

	logger.Log("load config:", viper.ConfigFileUsed())
	logger.Log("InitConfig Suc=", Global_Config)
	InitNowServerItem(http_port)

	return nil
}

func GetDB() DBConfig {
	return Global_Config.DB
}

/*
*
user:password        账号密码
@tcp(127.0.0.1:3306) 连接地址
/testdb              使用的数据库
?charset=utf8mb4     字符集		utf8mb4 = 支持中文 + emoji
&parseTime=true      时间解析	时间字段自动转成 Go 的 time.Time
&loc=Local           时区		用服务器本地时区
*/
func GetDBDns(dbConfig DBConfig, dbName string) string {
	// dsn := "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=true&loc=Local"
	// dsn := "user:password@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbName + "?charset=utf8mb4&parseTime=true&loc=Local"
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbName + "?charset=utf8mb4&parseTime=true&loc=Local"
	return dsn
}

func RandomServerItem(serverType string) ServerItem {
	var list []ServerItem
	switch serverType {
	case ServerType_List.Account_server:
		list = Global_Config.AccountServer
	case ServerType_List.Hall_server:
		list = Global_Config.HallServer
	}
	if len(list) == 0 {
		logger.Warn("RandomServerItem list is empty", serverType)
		return ServerItem{}
	}
	return randomServerItemByList(list)
}

func randomServerItemByList(list []ServerItem) ServerItem {
	if len(list) == 0 {
		logger.Warn("randomServerItemByList list is empty")
		return ServerItem{}
	}
	ru := random_util.NewRandomUtilAuto()
	return random_util.RandomItem(ru, list)
}
