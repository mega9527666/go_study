package string_util

import "mega/engine/logger"

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
