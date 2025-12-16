package string_util

import (
	"mega/engine/logger"
	"strconv"
	"time"
)

func GetStringFromMap(m map[string]interface{}, key string) (string, bool) {
	v, ok := m[key]
	if !ok {
		logger.Warn("GetStringFromMap no key error", key, m)
		return "", false
	}
	s, ok := v.(string)
	if !ok {
		logger.Warn("GetStringFromMap no string error", key, m)
	}
	return s, ok
}

// 当前毫秒级时间戳（string）🔥最常用	"1734245892123"
func NowUnixMilliString() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}
