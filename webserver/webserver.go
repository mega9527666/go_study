package main

import (
	"mega/engine/logger"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("我是indexHandler=", w)
}

func main() {
	logger.Log("webserver.main")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9090", nil)
}
