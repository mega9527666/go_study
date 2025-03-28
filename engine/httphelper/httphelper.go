package httphelper

import (
	"io"
	"mega/engine/logger"
	"net/http"
	"strings"
)

func Post(url string) {
	url = "http://127.0.0.1:15000/test?gameId=8901"
	jsonData := `{"title": "foo", "body": "bar", "userId": 1}`

	resp, err := http.Post(url, "application/json", strings.NewReader(jsonData))
	if err != nil {
		logger.Log("请求失败:", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log("读取响应失败:", err)
		return
	}

	logger.Log("响应:", string(body))
	defer resp.Body.Close()

}

func Get() {
	logger.Log("请求Get:")
	resp, err := http.Get("http://127.0.0.1:15000/test?gameId=8901")
	if err != nil {
		logger.Log("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log("读取响应失败:", err)
		return
	}

	logger.Log("响应:", string(body))
}
