package json_util

import (
	"encoding/json"
	"mega/engine/logger"
)

// Stringify Go 对象 -> JSON string
func Stringify(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		logger.Warn("json_util.Stringify error=", err)
		return ""
	}
	return string(b)
}

// Parse JSON string -> Go 对象
func Parse(jsonStr string, v any) error {
	err := json.Unmarshal([]byte(jsonStr), v)
	if err != nil {
		logger.Warn("json_util.Parse error=", err)
	}
	return err
}
