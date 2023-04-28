package smb

import (
	"fmt"
	"function/smb/smb/smb"
)

type SMB struct {
	Host     string
	Port     int64
	Proxy    string
	Username string // 用户名
	Password string // 密码
	Cmd      string
}

func (s *SMB) Smbconnect() {
	option := smb.Options{
		Host:     s.Host,
		Port:     int(s.Port),
		User:     s.Username,
		Password: s.Password,
	}
	if option.Proxy != "" {
		option.Proxy = s.Proxy
	}
	session, err := smb.NewSession(option, false)
	session.Close()
	if err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println("连接成功")
	}
}
