package handlers

import "net/http"

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	//允许跨源
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
	// 建议增加OPTIONS处理
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
