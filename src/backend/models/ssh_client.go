package models

type SshClient struct {
	Id           string `json:"id"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	UserName     string `json:"userName"`
	AuthType     string `json:"authType"`
	Password     string `json:"password"`
	Timeout      int    `json:"timeout"`
	Encoding     string `json:"encoding"`
	TerminalType string `json:"terminalType"`
}
