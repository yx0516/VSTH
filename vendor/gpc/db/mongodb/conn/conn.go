package conn

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

//-----------------------------------------------------------------------------------------------------------//

// 登录后的 Session
type Session struct {
	*mgo.Session
}

func NewSession(session *mgo.Session) *Session {
	return &Session{session}
}

// 退出登录并且关闭连接
func (self *Session) SafeClose() {
	if self.Session != nil {
		self.Session.LogoutAll()
		self.Session.Close()
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 连接配置
type Conn struct {
	IP       string
	Port     int
	UserName string
	UserPwd  string
	DB       string
	Timeout  int // 超时，单位：秒，默认10秒
}

func NewConn(ip string, port int) *Conn {
	return &Conn{
		IP:   ip,
		Port: port,
	}
}

// 检查字段
func (self *Conn) Check() *Conn {
	self.IP = strings.TrimSpace(self.IP)

	if self.Port <= 0 {
		self.Port = 27017
	}

	self.UserName = strings.TrimSpace(self.UserName)
	self.UserPwd = strings.TrimSpace(self.UserPwd)
	self.DB = strings.TrimSpace(self.DB)

	if self.Timeout <= 0 {
		self.Timeout = 10
	}

	return self
}

// 登录成功后返回 Session
func (self *Conn) Login() (*Session, error) {
	self.Check()
	host := fmt.Sprintf("%v:%v", self.IP, self.Port)

	if session, err := mgo.DialWithTimeout(host, time.Second*time.Duration(self.Timeout)); err != nil {
		return nil, err
	} else {
		cred := &mgo.Credential{ // 凭证
			Username: self.UserName,
			Password: self.UserPwd,
			Source:   self.DB,
		}
		if err = session.Login(cred); err != nil {
			return nil, err
		} else {
			return NewSession(session), nil
		}
	}
}

//-----------------------------------------------------------------------------------------------------------//
