package dateutil

import (
	"strconv"
	"time"
)

func init() {
	// fmt.Println("dateutil.init")
}

// 打印日志的时间
func FormatLogNow() string {
	var currentTime = time.Now()
	var timestamp = currentTime.UnixMilli()
	return FormatTime_2(timestamp) + " " + strconv.FormatInt(timestamp, 10)
}

// 按要求格式化时间戳
func FormatTime(timestamp int64) string {
	// 第二个参数是纳秒，通常为0
	var tempTime = time.Unix(timestamp, 0)
	var formatTime = tempTime.Format("2006-01-02 15:04:05")
	return formatTime
}

// 按要求格式化时间戳
func FormatTime_2(timestamp int64) string {
	// 第二个参数是纳秒，通常为0
	var tempTime = time.Unix(timestamp, 0)
	var formatTime = tempTime.Format("2006-01-02 15:04:05.000")
	return formatTime
}

// 返回毫秒的时间戳
func Now_UnixMilli() int64 {
	return time.Now().UnixMilli()
}

func Now_Unix() int64 {
	return time.Now().Unix()
}

func Now_UnixMicro() int64 {
	return time.Now().UnixMicro()
}
