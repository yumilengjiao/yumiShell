package handlers

import (
  "backend/config"
  "context"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"
  "github.com/pkg/sftp"
  "golang.org/x/crypto/ssh"
  "log"
  "net/http"
  "strconv"
)

// 用于存储所有的websocket连接
var (
  allsshConns map[string]*sshConn
)

func init() {
  //初始化存储连接的map
  allsshConns = make(map[string]*sshConn)
}

// ssh客户端结构体
type sshConn struct {
  conn       *ssh.Client
  sshSession *ssh.Session
  sftpClient *sftp.Client
}

//type result struct {
//	terminalMessage []byte
//	fileList        []string
//}

// HandleWebsocket 用来处理客户端连接
func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
  upgrade := initWebsocket(w, r)
  //获取路径参数
  vars := mux.Vars(r)
  //在session配置文件中查找配置
  sessionConfig := config.GetSession(vars["id"])
  //获取每个前端连接的唯一id，用于前端界面尺寸变化的时候调整所有终端的尺寸
  uniqId := vars["uniqId"]

  log.Println("传入的路径参数是", vars, "本次连接的session是", sessionConfig)
  //初始化ssh客户端结构体
  sshClient := initSshConn(sessionConfig)
  defer sshClient.close()
  //存储ssh连接
  allsshConns[uniqId] = sshClient
  //开始进行终端交互
  err := sshClient.terminalInteraction(upgrade, sessionConfig)
  if err != nil {
    log.Println("终端交互失败")
    return
  }
  log.Println("终端交互结束")
  return
}

// 初始化ssh客户端结构体的方法
func initSshConn(sessionConfig *config.Session) (sshClient *sshConn) {
  sshClient = &sshConn{}
  sshConfig := &ssh.ClientConfig{
    User: sessionConfig.Username,
    Auth: []ssh.AuthMethod{
      //密码验证
      ssh.Password(sessionConfig.Password),
    },
    HostKeyCallback: ssh.InsecureIgnoreHostKey(),
  }
  conn, err := ssh.Dial("tcp", sessionConfig.Host+":"+strconv.Itoa(sessionConfig.Port), sshConfig)
  if err != nil {
    log.Println("ssh连接失败", err)
  }
  //sshClient值初始化
  sshClient.sshSession, err = conn.NewSession()
  sshClient.conn = conn
  sshClient.sftpClient, err = sftp.NewClient(conn)
  return
}

// 初始化websocket
func initWebsocket(w http.ResponseWriter, r *http.Request) (upgrade *websocket.Conn) {
  //初始化websocket变量
  upgrader := websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
  }
  upgrade, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println("websocket升级失败", err)
  }
  return
}

// 进行终端交互
func (sshClient *sshConn) terminalInteraction(upgrade *websocket.Conn, sshConfig *config.Session) error {
  ctx := context.Background()
  defer ctx.Done()
  //设置终端类型
  err := sshClient.sshSession.RequestPty(sshConfig.TerminalType, 80, 40, ssh.TerminalModes{
    ssh.ECHO:          0,     // 不启用回显
    ssh.TTY_OP_ISPEED: 14400, // 输入速度 = 14.4kbaud
    ssh.TTY_OP_OSPEED: 14400, // 输出速度 = 14.4kbaud
  })
  if err != nil {
    log.Println("设置终端类型失败", err)
    return err
  }
  // 获取标准输入输出流
  stdin, err := sshClient.sshSession.StdinPipe()
  if err != nil {
    log.Println("获取stdin失败", err)
    return err
  }
  defer stdin.Close()

  stdout, err := sshClient.sshSession.StdoutPipe()
  if err != nil {
    log.Println("获取stdout失败", err)
    return err
  }

  stderr, err := sshClient.sshSession.StderrPipe()
  if err != nil {
    log.Println("获取stderr失败", err)
    return err
  }
  // 启动交互式shell
  err = sshClient.sshSession.Shell()
  if err != nil {
    log.Println("开启shell失败", err)
    return err
  }
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  // 创建goroutine处理stdout输出
  go func() {
    defer cancel()
    buffer := make([]byte, 1024)
    for {
      n, err := stdout.Read(buffer)
      if n > 0 {
        log.Println("stdout数据:", string(buffer[:n]))
        if n == 0 {
          continue
        }
        upgrade.WriteMessage(websocket.TextMessage, buffer[:n])
      }
      if err != nil {
        log.Println("stdout读取错误:", err)
        return
      }
    }
  }()

  // 创建goroutine处理stderr输出
  go func() {
    defer cancel()
    buffer := make([]byte, 1024)
    for {
      n, err := stderr.Read(buffer)
      if n > 0 {
        log.Println("stderr数据:", string(buffer[:n]))
        upgrade.WriteMessage(websocket.TextMessage, buffer[:n])
      }
      if err != nil {
        log.Println("stderr读取错误:", err)
        return
      }
    }
  }()

  // 主循环处理websocket输入
  for {
    select {
    case <-ctx.Done():
      return nil
    default:
      _, p, err := upgrade.ReadMessage()
      if err != nil {
        log.Println("websocket读取失败", err)
        return err
      }
      log.Println("收到websocket数据:", p)
      // 移除非打印字符（保留 ASCII 可打印字符和 \n）
      cleaned := make([]byte, 0, len(p))
      for _, b := range p {
        if b >= 0x20 && b <= 0x7E || b == '\n' || b == '\r' { // 保留空格到~和换行符
          cleaned = append(cleaned, b)
        }
      }
      log.Printf("清理后的命令: %q", string(cleaned))

      // 确保命令以 \n 结尾
      if len(cleaned) == 0 || cleaned[len(cleaned)-1] != '\n' {
        cleaned = append(cleaned, '\n')
      }
      _, err = stdin.Write(cleaned)
      if err != nil {
        log.Println("写入stdin失败", err)
        return err
      }
    }
  }
  return nil
}

// 关闭ssh客户端连接的方法
func (sshClient *sshConn) close() {
  err1 := sshClient.sshSession.Close()
  if err1 != nil {
    return
  }
  err2 := sshClient.conn.Close()
  if err2 != nil {
    return
  }
  err3 := sshClient.sftpClient.Close()
  if err3 != nil {
    return
  }
}
