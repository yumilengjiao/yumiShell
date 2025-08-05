package main

import (
	"backend/config"
	"backend/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func main() {
	log.Println(config.SessionGroups)
	// 启动 HTTP 服务器
	router.HandleFunc("/ws/term/{id}/{uniqId}", handlers.HandleWebsocket)

	http.ListenAndServe(":8080", router)
}
