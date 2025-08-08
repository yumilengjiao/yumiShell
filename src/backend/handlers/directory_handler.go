package handlers

import (
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
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
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
    log.Println("接收的目录文件路径是", string(bytes))
    path := string(bytes)
    //发送当前路径
    err = upgrade.WriteMessage(websocket.TextMessage, []byte(path))
    if err != nil {
      log.Println("发送当前路径失败", err)
      continue
    }

    //第一次直接发送当前路径下的所有文件
    dir, err := client.sftpClient.ReadDir(path)
    if err != nil {
      log.Println("获取目录文件失败", err)
      continue
    }
    var fileList []FileItem
    for _, file := range dir {
      fileList = append(fileList, FileItem{
        IsDir:   file.IsDir(),
        Name:    file.Name(),
        Size:    file.Size(),
        Mode:    file.Mode().String(),
        ModTime: time.Time.Format(file.ModTime(), "2006-01-02 15:04:05"),
      })
    }
    //返回目录文件
    err = upgrade.WriteJSON(fileList)
    if err != nil {
      log.Println("返回目录文件失败", err)
      continue
    }
  }
}
