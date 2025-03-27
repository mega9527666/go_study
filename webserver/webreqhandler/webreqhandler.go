package webreqhandler

import (
	"mega/engine/logger"
	"net/http"
	"os"
	"strconv"
)

func init() {
	logger.Log("webreqhandler.init")
}

func megaHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("megaHandler=Host=", r.RequestURI, r.Host, r.RemoteAddr)
}

func Init(port int) {
	var portStr string = strconv.Itoa(port)
	logger.Log("webreqhandler.Init")
	var dirPath string = "public/public" + portStr
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			logger.Log("创建文件夹失败==", dirPath, err)
		} else {
			logger.Log("创建文件夹成功==", dirPath)
		}
	} else {
		logger.Log("已经存在文件夹==", dirPath)
	}

	// 设置静态文件服务器的根目录为当前目录
	fs := http.FileServer(http.Dir(dirPath))
	// 设置路由规则，将所有请求重定向到静态文件服务器
	http.Handle("/", fs)
	http.HandleFunc("/mega", megaHandler)
	http.ListenAndServe(":"+portStr, nil)

}
