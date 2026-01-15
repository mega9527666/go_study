package main

import (
	"encoding/json"
	"mega/engine/http_client"
	"mega/engine/json_util"
	"mega/engine/logger"
	"mega/engine/md5_helper"
)

// "github.com/gorilla/websocket"

func main() {

	dataObj := map[string]any{
		"v": "1.0.2",
		"t": 123456789,
	}
	dataStr := json_util.Stringify(dataObj)
	realParam := map[string]string{
		"k":    md5_helper.GetMd5_encrypt(dataStr),
		"data": dataStr,
	}

	client := http_client.NewHttpClient(0)
	// GET
	resp, err := client.Post(
		// "https://ipinfo.io/json",
		"http://127.0.0.1:9090/init_client",
		realParam,
		map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
	)

	if err != nil {
		logger.Warn("请求错误=", err)
	} else {
		// logger.Log("请求成功=", resp)
		var dataObj map[string]interface{}
		if err := json.Unmarshal([]byte(resp), &dataObj); err != nil {

		} else {
			logger.Log("请求成功=json--", dataObj)
		}

	}
}
