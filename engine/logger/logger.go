package logger

import (
	"fmt"
	"mega/engine/dateutil"
)

const (
	LOG_LEVEL_DEBUG = iota
	LOG_LEVEL_LOG
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
)

const TAG = "MegaLog"

var Log_Level_Names = [5]string{"debug", "log", "info", "warn", "error"}
var LOG_LEVEL = LOG_LEVEL_DEBUG

func init() {
	// fmt.Println("logger.init")
}

func getLogPreKey(nowLevel int) string {
	var str string = "[" + dateutil.FormatNow() + "] " + TAG + " [" + Log_Level_Names[nowLevel] + "] "
	return str
}

func Debug(a ...any) {
	if LOG_LEVEL > LOG_LEVEL_DEBUG {
		return
	}
	var str string = getLogPreKey(LOG_LEVEL_DEBUG)
	fmt.Print(str)
	fmt.Println(a...)
}

func Log(a ...any) {
	if LOG_LEVEL > LOG_LEVEL_LOG {
		return
	}
	var str string = getLogPreKey(LOG_LEVEL_LOG)
	fmt.Print(str)
	fmt.Println(a...)
}

func Info(a ...any) {
	if LOG_LEVEL > LOG_LEVEL_INFO {
		return
	}
	var str string = getLogPreKey(LOG_LEVEL_INFO)
	fmt.Print(str)
	fmt.Println(a...)
}

func Warn(a ...any) {
	if LOG_LEVEL > LOG_LEVEL_WARN {
		return
	}
	var str string = getLogPreKey(LOG_LEVEL_WARN)
	fmt.Print(str)
	fmt.Println(a...)
}

func Error(a ...any) {
	if LOG_LEVEL > LOG_LEVEL_ERROR {
		return
	}
	var str string = getLogPreKey(LOG_LEVEL_ERROR)
	fmt.Print(str)
	fmt.Println(a...)
}
