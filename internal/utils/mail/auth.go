package mail

import (
	"com.caiflower/commons/thirdpart/internal/config"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/smtp"
)

type (
	LoginAuth struct {
		username string
		password string
	}
)

func NewAuthLogin(username, password string) *LoginAuth {
	return &LoginAuth{username: username, password: password}
}

func (auth *LoginAuth) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	return "LOGIN", []byte{}, nil
}

func (auth *LoginAuth) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	if more {
		switch string(fromServer) {
		case "username:":
			return []byte(auth.username), nil
		case "Password:":
			return []byte(auth.password), nil
		default:
			return nil, errors.New("unknown fromServer")
		}
	}
	return nil, nil
}

func (auth *LoginAuth) GetClient() (client *Client, e error) {
	address := config.Config.Address
	host, _, _ := net.SplitHostPort(address)
	// 1.连接25端口smtp
	c, e := smtp.Dial(address)
	if e != nil {
		log.Println(e)
		return
	}

	// 2.建立ssl
	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		if e = c.StartTLS(config); e != nil {
			log.Println("mail tls connect err")
		}
	}

	// 3.测试认证信息
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if e = c.Auth(auth); e != nil {
				log.Printf("mail tls connect err:%v\n", e)
				return
			}
		}
	}
	client = &Client{client: c}
	return
}
