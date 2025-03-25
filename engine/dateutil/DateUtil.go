package DateUtil

import (
	"fmt"
	"time"
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
	fmt.Println("DateUtil.init")
}

func FormatNow() string {
	var currentTime = time.Now()
	var formatTime = currentTime.Format("2006-01-02 15:04:05")
	// fmt.Println("formatNow==Now", currentTime)
	// fmt.Println("formatNow==Unix", currentTime.Unix())
	fmt.Println("formatNow==formatTime", formatTime)
	return formatTime
}

func FormatTime(timestamp int64) string {
	// 第二个参数是纳秒，通常为0
	var tempTime = time.Unix(timestamp, 0)
	var formatTime = tempTime.Format("2006-01-02 15:04:05")
	// fmt.Println("FormatTime==tempTime", tempTime)
	fmt.Println("FormatTime==formatTime", formatTime)
	return ""
}
