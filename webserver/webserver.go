package main

import (
	"mega/engine/logger"
	"net/http"
	// "github.com/gorilla/websocket"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("我是indexHandler=Host=", r.RequestURI, r.Host, r.RemoteAddr)
}

func megaHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("megaHandler=Host=", r.RequestURI, r.Host, r.RemoteAddr)
}

func main() {
	logger.Log("webserver.main")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/mega", megaHandler)
	http.ListenAndServe(":9090", nil)
}
