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
	router.HandleFunc("/ws/term/{id}/{uniqId}", handlers.HandleTerminal)
	router.HandleFunc("/ws/file/{uniqId}", handlers.HandleDirectory)

	err := http.ListenAndServe(":8977", router)
	if err != nil {
		return
	}
}
