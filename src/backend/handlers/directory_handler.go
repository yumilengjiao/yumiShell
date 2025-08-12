package handlers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sort"
	"time"
)

type FileItem struct {
	IsDir   bool   `json:"isDir"`
	ModTime string `json:"modTime"`
	Mode    string `json:"mode"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
}
type SftpRequest struct {
	RequestType string `json:"requestType"`
	Path        string `json:"path"`
	// 文件内容（用于上传操作）
	File struct {
		Name    string `json:"name"`    // 文件名
		Content string `json:"content"` // Base64编码的文件内容
	}
	// 是否递归操作（用于目录）
	Recursive bool `json:"recursive"`
}

// HandleDirectory Hand 处理目录请求
func HandleDirectory(w http.ResponseWriter, r *http.Request) {
	//获取当前的唯一id
	vars := mux.Vars(r)
	s := vars["uniqId"]
	log.Println("传入的uniq参数是", s)
	//获取当前路径的sshClient
	mutex.Lock()
	client := sshClientList[s]
	mutex.Unlock()
	upgrader := websocket.Upgrader{
		ReadBufferSize:  128 * 1024,
		WriteBufferSize: 128 * 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	upgrade, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket升级失败", err)
	}

	//当前路径下的文件列表
	var fileList []FileItem
	//获取当前目录
	initialPath, err := client.sftpClient.Getwd()
	if err != nil {
		log.Println("获取当前目录失败", err)
	}
	log.Println("当前目录是", initialPath)
	//发送初始化路径
	err = upgrade.WriteMessage(websocket.TextMessage, []byte(initialPath))
	//获取当前目录下的所有文件
	dir, err := client.sftpClient.ReadDir(initialPath)
	if err != nil {
		log.Println("获取当前目录失败", err)
	}
	sort.Slice(dir, func(i, j int) bool {
		return dir[i].Name() < dir[j].Name()
	})
	// 手动添加 "." 和 ".." 项
	fileList = append(fileList, FileItem{
		IsDir:   true,
		Name:    ".",
		Size:    0,
		Mode:    "drwxr-xr-x", // 模拟目录权限
		ModTime: time.Now().Format("2006-01-02 15:04:05"),
	})

	// 如果不是根目录，添加 ".." 项
	if initialPath != "/" {
		fileList = append(fileList, FileItem{
			IsDir:   true,
			Name:    "..",
			Size:    0,
			Mode:    "drwxr-xr-x", // 模拟目录权限
			ModTime: time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	for _, file := range dir {
		fileList = append(fileList, FileItem{
			IsDir:   file.IsDir(),
			Name:    file.Name(),
			Size:    file.Size(),
			Mode:    file.Mode().String(),
			ModTime: time.Time.Format(file.ModTime(), "2006-01-02 15:04:05"),
		})
	}
	//websocket连接时发送当前目录信息
	log.Println("返回的fileList:", fileList)
	err = upgrade.WriteJSON(fileList)
	if err != nil {
		return
	}
	go handleDirectoryFile(upgrade, client)
}

// 处理目录文件请求
func handleDirectoryFile(upgrade *websocket.Conn, client *sshClient) {
	defer upgrade.Close()
	for {
		_, bytes, err := upgrade.ReadMessage()
		if err != nil {
			log.Println("读取目录文件失败", err)
			return
		}
		//解析前端传输类型
		var sftpRequest SftpRequest
		err = json.Unmarshal(bytes, &sftpRequest)
		if err != nil {
			log.Println("解析目录文件失败", err)
			continue
		}
		//通过前端传输不同类型调用不同方法
		switch sftpRequest.RequestType {
		case "get":
			err := toGetPathFile(&sftpRequest, client, upgrade)
			if err != nil {
				continue
			}
		case "upload":
			err := toUploadFile(&sftpRequest, client)
			if err != nil {
				continue
			}
		case "download":
			err := toDownloadFile(&sftpRequest, client, upgrade)
			if err != nil {
				continue
			}
		case "delete":
			err := toDeleteFile(&sftpRequest, client, upgrade)
			if err != nil {
				continue
			}
		case "mkdir":
			err := toCreateDirectory(&sftpRequest, client, upgrade)
			if err != nil {
				continue
			}
		case "create":
			err := toCreateFile(&sftpRequest, client, upgrade)
			if err != nil {
				continue
			}
		case "rmDir":
			err := toDeleteDirectory(&sftpRequest, client, upgrade)
			if err != nil {
				continue
			}
		}
	}
}
func toGetPathFile(sftpRequest *SftpRequest, client *sshClient, upgrade *websocket.Conn) error {
	log.Println("获取路径下的所有文件")
	var fileList []FileItem
	dir, err := client.sftpClient.ReadDir(sftpRequest.Path)
	if err != nil {
		log.Println("获取当前目录失败", err)
	}
	sort.Slice(dir, func(i, j int) bool {
		return dir[i].Name() < dir[j].Name()
	})

	// 手动添加 "." 和 ".." 项
	fileList = append(fileList, FileItem{
		IsDir:   true,
		Name:    ".",
		Size:    0,
		Mode:    "drwxr-xr-x", // 模拟目录权限
		ModTime: time.Now().Format("2006-01-02 15:04:05"),
	})

	// 如果不是根目录，添加 ".." 项
	if sftpRequest.Path != "/" {
		fileList = append(fileList, FileItem{
			IsDir:   true,
			Name:    "..",
			Size:    0,
			Mode:    "drwxr-xr-x", // 模拟目录权限
			ModTime: time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	for _, file := range dir {
		fileList = append(fileList, FileItem{
			IsDir:   file.IsDir(),
			Name:    file.Name(),
			Size:    file.Size(),
			Mode:    file.Mode().String(),
			ModTime: time.Time.Format(file.ModTime(), "2006-01-02 15:04:05"),
		})
	}
	//websocket连接时发送当前目录信息
	log.Println("返回的fileList:", fileList)
	err = upgrade.WriteJSON(fileList)
	if err != nil {
		log.Println("websocket发送数据失败", err)
		return err
	}
	return nil
}
func toUploadFile(sftpRequest *SftpRequest, client *sshClient) error {
	log.Println("上传文件")
	//获取上传路径
	path := sftpRequest.Path
	log.Println("上传路径是", path)
	log.Println("上传的文件名是", sftpRequest.File.Name)
	//base64转码文件
	file, err := base64.StdEncoding.DecodeString(sftpRequest.File.Content)
	if err != nil {
		log.Println("base64解码失败", err)
		return err
	}
	create, err := client.sftpClient.Create(path + "/" + sftpRequest.File.Name)
	log.Println("创建文件的路径为", path+"/"+sftpRequest.File.Name)
	if err != nil {
		log.Println("创建文件失败", err)
		return err
	}
	_, err = create.Write(file)
	if err != nil {
		log.Println("写入文件失败", err)
	}
	return nil
}
func toDownloadFile(sftpRequest *SftpRequest, client *sshClient, upgrade *websocket.Conn) error {
	log.Println("下载文件")
	return nil
}
func toDeleteFile(sftpRequest *SftpRequest, client *sshClient, upgrade *websocket.Conn) error {
	log.Println("删除文件")
	return nil
}
func toCreateFile(sftpRequest *SftpRequest, client *sshClient, upgrade *websocket.Conn) error {
	log.Println("创建文件")
	return nil
}
func toCreateDirectory(sftpRequest *SftpRequest, client *sshClient, upgrade *websocket.Conn) error {
	log.Println("创建目录")
	return nil
}
func toDeleteDirectory(sftpRequest *SftpRequest, client *sshClient, upgrade *websocket.Conn) error {
	log.Println("删除目录")
	return nil
}
