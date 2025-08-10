package handlers

import (
	"backend/config"
	"context"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	mutex         sync.Mutex
	sshClientList = make(map[string]*sshClient)
)

type sshClient struct {
	clientConfig  *ssh.ClientConfig
	client        *ssh.Client
	session       *ssh.Session
	sessionConfig *config.Session
	sftpClient    *sftp.Client
	stdin         io.WriteCloser
	stdout        io.Reader
	stderr        io.Reader
}

// HandleTerminal HandleWebsocket 处理前端websocket请求的方法
func HandleTerminal(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	//初始化websocket
	upgrader := websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
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
	//获取路径参数
	vars := mux.Vars(r)
	sessionId := vars["id"]
	log.Println("sessionId是", sessionId)
	uniqId := vars["uniqId"]
	log.Println("uniqId是", uniqId)
	//获取session配置对象
	sessionConfig := config.GetSession(sessionId)

	//初始化ssh客户端结构体
	client, err := initSshClient(sessionConfig)
	if err != nil {
		log.Println("sshClient初始化失败", err)
	}
	mutex.Unlock()
	defer client.close()
	//全局中存储sshClient对象
	sshClientList[uniqId] = client
	//开启ssh并处理数据和websocket
	err = client.start(upgrade)
	if err != nil {
		log.Println("sshClient启动失败", err)
		return
	}
}

// 初始化sshClient
func initSshClient(session *config.Session) (*sshClient, error) {
	sshConfig := ssh.ClientConfig{
		User: session.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(session.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	address := session.Host + ":" + strconv.Itoa(session.Port)
	// 建立ssh连接
	conn, err := ssh.Dial("tcp", address, &sshConfig)
	if err != nil {
		return nil, err
	}
	// 初始化ssh session
	newSession, err := conn.NewSession()
	//初始化sftp
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		err1 := conn.Close()
		if err1 != nil {
			return nil, err1
		}
		return nil, err1
	}
	//设置pty模式
	newSession.RequestPty(session.TerminalType, 34, 90, ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	})
	//获取连接管道
	stdin, err := newSession.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := newSession.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := newSession.StderrPipe()
	if err != nil {
		return nil, err
	}
	return &sshClient{
		clientConfig:  &sshConfig,
		client:        conn,
		session:       newSession,
		sessionConfig: session,
		sftpClient:    sftpClient,
		stdin:         stdin,
		stdout:        stdout,
		stderr:        stderr,
	}, nil
}

// 开启ssh连接
func (sshClient *sshClient) start(upgrade *websocket.Conn) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := sshClient.session.Shell()
	if err != nil {
		return err
	}
	// 处理stdout数据
	go func() {
		defer cancel()
		buffer := make([]byte, 1024)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n, err := sshClient.stdout.Read(buffer)
				if err != nil {
					log.Println("stdout读取错误:", err)
					return
				}
				log.Println("stdout数据:", string(buffer[:n]))
				upgrade.WriteMessage(websocket.TextMessage, buffer[:n])
			}
		}
	}()
	// 处理stderr数据
	go func() {
		defer cancel()
		buffer := make([]byte, 1024)
		select {
		case <-ctx.Done():
			return
		default:
			n, err := sshClient.stderr.Read(buffer)
			if err != nil {
				log.Println("stderr读取错误:", err)
				return
			}
			log.Println("元数据是", buffer[:n])
			log.Println("stderr数据:", string(buffer[:n]))
			upgrade.WriteMessage(websocket.TextMessage, buffer[:n])
		}
	}()
	//处理stdin的值
	for {
		_, p, err := upgrade.ReadMessage()
		log.Println("收到前端发送数据", string(p))
		if err != nil {
			log.Println("websocket读取失败", err)
			return err
		}
		if strings.Contains(string(p), "cd ") {
			log.Println("打印的命令是", string(p))
		}
		_, err = sshClient.stdin.Write(p)
		if err != nil {
			log.Println("websocket写入失败", err)
			return err
		}

	}
}

// 关闭sshClient
func (sshClient *sshClient) close() {
	err := sshClient.session.Close()
	if err != nil {
		log.Println("sshClient关闭失败", err)
	}
	err = sshClient.client.Close()
	if err != nil {
		log.Println("sshClient关闭失败", err)
	}
	err = sshClient.stdin.Close()
	if err != nil {
		log.Println("sshClient关闭失败", err)
	}
}
