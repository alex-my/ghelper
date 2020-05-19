package email

import (
	"strings"

	"github.com/go-gomail/gomail"
)

// Config 邮件配置
type Config struct {
	// Host SMTP server 地址, 例如 smtp.example.com
	Host string
	// Port SMTP server 端口, 例如 25, 465, 587
	Port int
	// Username 邮箱账号
	Username string
	// Password 邮箱密码或者授权码
	Password string
}

var config *Config

// SetConfig 初始化参数
// 服务器启动时初始化一次
func SetConfig(c *Config) {
	config = c
}

// Send 发送邮件
// from: 发送人地址，可以是 xxx@xxx，也可以是 XXX <xxx@xxx>
// to: 目标邮箱地址，多个地址使用英文逗号隔离
// subject: 邮件标题
// html: 邮件内容, 支持 text/html 格式
func Send(from, to, subject, html string) error {
	m := gomail.NewMessage()

	toers := []string{}
	for _, address := range strings.Split(to, ",") {
		toers = append(toers, address)
	}

	m.SetHeader("From", from)
	m.SetHeader("To", toers...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	return d.DialAndSend(m)
}
