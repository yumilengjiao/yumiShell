package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	SessionGroups []SessionGroup
	InitSignal    = make(chan struct{})
)

func init() {
	LoadSessions()
	close(InitSignal)
}

type Session struct {
	Id           string `json:"id"`
	LabelName    string `json:"labelName"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	AuthType     string `json:"authType"`
	Password     string `json:"password"`
	PrivateKey   string `json:"privateKey"`
	Passphrase   string `json:"passphrase"`
	TimeOut      int    `json:"timeOut"`
	Encoding     string `json:"encoding"`
	TerminalType string `json:"terminalType"`
}
type SessionGroup struct {
	Id          string     `json:"id"`
	Label       string     `json:"label"`
	SessionList []*Session `json:"sessionList"`
}

// LoadSessions 加载session.json文件
func LoadSessions() (complete bool) {
	SessionGroups = []SessionGroup{}
	log.Println("Load方法开始执行")
	//获取AppData路径
	appDataPath, err := os.UserConfigDir()
	if err != nil {
		log.Println("获取AppData路径失败")
	}
	sessionPath := filepath.Join(appDataPath, "yumi-shell", "sessions.json")
	if err != nil {
		log.Println(err)
	}
	//session.json文件读取
	file, err := os.ReadFile(sessionPath)
	if err != nil {
		log.Println("session.json文件不存在")
		return
	}
	time.Sleep(5 * time.Second)
	log.Println("session.json文件读取成功，读取的文件内容是", string(file))
	err = json.Unmarshal(file, &SessionGroups)
	if err != nil {
		log.Println("session.json文件解析失败", err)
		return
	}
	log.Println("解析后的SessionGroups", SessionGroups)
	log.Println("session.json文件解析成功")
	return true
}

// GetSession 获得指定session
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
