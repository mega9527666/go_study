package logger

import (
	"fmt"
)

const (
	LOG_LEVEL_DEBUG = iota
	LOG_LEVEL_LOG
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
)

const TAG = "MegaLog"

var LOG_LEVEL = LOG_LEVEL_DEBUG

func init() {
	fmt.Println("logger.init")
}

// func getLogPreKey() string {
// 	// let str: string = "[" + DateUtil.formatNow() + "] " + Logger.tag + " ["
// 	// 	+ Log_Level_Names[nowLevel] + "] ";
// 	// return str;
// 	return ""
// }

func Log(a ...any) (n int, err error) {
	if LOG_LEVEL > LOG_LEVEL_LOG {
		return
	}
	return fmt.Println(a...)
}
