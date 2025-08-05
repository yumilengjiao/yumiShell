package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var (
	SessionGroups = make([]SessionGroup, 100)
)

func init() {
	LoadSessions()
}

type Session struct {
	Id           string `json:"id"`
	LabelName    string `json:"labelName"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	AuthType     string `json:"authType"`
	Password     string `json:"password"`
	TimeOut      int    `json:"timeOut"`
	Encoding     string `json:"encoding"`
	TerminalType string `json:"terminalType"`
}
type SessionGroup struct {
	Id          string     `json:"id"`
	Label       string     `json:"label"`
	SessionList []*Session `json:"sessionList"`
}

// 加载session.json文件
func LoadSessions() {
	//获取AppData路径
	appDataPath, err := os.UserConfigDir()
	if err != nil {
		log.Println("获取AppData路径失败")
	}
	sessionPath := filepath.Join(appDataPath, "yumi-shell", "sessions.json")
	if err != nil {
		log.Println(err)
	}
	log.Println("AppData的路径是", sessionPath)
	//session.json文件读取
	file, err := os.ReadFile(sessionPath)
	if err != nil {
		log.Println("session.json文件不存在")
		return
	}
	err = json.Unmarshal(file, &SessionGroups)
	if err != nil {
		log.Println("session.json文件解析失败")
		return
	}
	log.Println("文件读取成功:session.json文件内容是", SessionGroups)
	log.Println("第一个会话组的session是", SessionGroups[0].SessionList[0])

}

// 获得指定session
func GetSession(sessionId string) *Session {
	for _, sessionGroup := range SessionGroups {
		for _, session := range sessionGroup.SessionList {
			if session.Id == sessionId {
				return session
			}
		}
	}
	return nil
}
