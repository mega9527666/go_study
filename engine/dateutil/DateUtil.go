package dateutil

import (
	"strconv"
	"time"
)

func init() {
	// fmt.Println("dateutil.init")
}

func FormatNow() string {
	var currentTime = time.Now()
	var formatTime = currentTime.Format("2006-01-02 15:04:05.000")
	return formatTime + " " + strconv.FormatInt(currentTime.UnixMilli(), 10)
}

func FormatTime(timestamp int64) string {
	// 第二个参数是纳秒，通常为0
	var tempTime = time.Unix(timestamp, 0)
	var formatTime = tempTime.Format("2006-01-02 15:04:05")
	return formatTime
}
