package handlers

import (
	"backend/config"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Content struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
type OtherRequest struct {
	ReqType        string                `json:"reqType"`
	ResizeContent  Content               `json:"content"`
	SessionContent []config.SessionGroup `json:"sessionContent"`
}

func OtherHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	upgrade, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket升级失败", err)
		return
	}
	defer upgrade.Close()
	go func() {
		<-config.InitSignal
		err2 := upgrade.WriteMessage(websocket.TextMessage, []byte("ok"))
		if err2 != nil {
			log.Println("websocket发送数据失败", err2)
		}
		log.Println("完成连接信号")
	}()
	//完成连接信号
	for {
		_, p, err := upgrade.ReadMessage()
		if err != nil {
			log.Println("websocket读取数据失败", err)
			return
		}
		log.Println("收到前端发送数据", string(p))
		request := &OtherRequest{}
		err = json.Unmarshal(p, request)
		if err != nil {
			log.Println("json解析失败", err)
			return
		}
		log.Println("请求类型是", request.ReqType)
		switch request.ReqType {
		case "fitness":
			resizeTerminal(request.ResizeContent.Row, request.ResizeContent.Col)
		case "readSessions":
			readSessions(upgrade, request.SessionContent)
		}

	}
}

// 重新调整所有终端的窗口大小
func resizeTerminal(row, col int) {
	log.Println("重新调整终端窗口大小", row, col)
	mutex.Lock()
	for _, value := range sshClientList {
		err := value.session.WindowChange(row, col)
		if err != nil {
			log.Println("窗口改变失败", err)
		}
	}
	mutex.Unlock()

}

// 重新加载session.json文件
func readSessions(upgrade *websocket.Conn, sessionContent []config.SessionGroup) {
	config.SessionGroups = sessionContent
}
