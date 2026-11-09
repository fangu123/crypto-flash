package util

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

// EmailConfig 邮件配置
type EmailConfig struct {
	Host     string // SMTP服务器地址
	Port     int    // SMTP端口
	User     string // 邮箱账号
	Password string // 邮箱密码或授权码
	From     string // 发件人地址
}

// Email 邮件发送器
type Email struct {
	config *EmailConfig
}

// NewEmail 创建邮件发送器实例
func NewEmail(config *EmailConfig) *Email {
	return &Email{
		config: config,
	}
}

// Send 发送邮件
func (e *Email) Send(to []string, subject, body string, isHTML bool) error {
	// 构建邮件内容
	contentType := "text/plain; charset=UTF-8"
	if isHTML {
		contentType = "text/html; charset=UTF-8"
	}

	msg := fmt.Sprintf("To: %s\r\n", strings.Join(to, ","))
	msg += fmt.Sprintf("From: %s\r\n", e.config.From)
	msg += fmt.Sprintf("Subject: %s\r\n", subject)
	msg += fmt.Sprintf("Content-Type: %s\r\n\r\n", contentType)
	msg += body

	// 构建SMTP服务器地址
	addr := fmt.Sprintf("%s:%d", e.config.Host, e.config.Port)

	// 创建SMTP客户端
	auth := smtp.PlainAuth("", e.config.User, e.config.Password, e.config.Host)

	// 根据端口判断是否使用TLS
	var err error
	if e.config.Port == 465 {
		// 使用SSL/TLS连接
		err = e.sendWithTLS(addr, auth, e.config.From, to, []byte(msg))
	} else {
		// 使用STARTTLS或普通连接
		err = smtp.SendMail(addr, auth, e.config.From, to, []byte(msg))
	}

	return err
}

// sendWithTLS 使用TLS发送邮件
func (e *Email) sendWithTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	// 创建TLS配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         e.config.Host,
	}

	// 建立TCP连接
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	// 创建SMTP客户端
	client, err := smtp.NewClient(conn, e.config.Host)
	if err != nil {
		return err
	}
	defer client.Close()

	// 认证
	if err = client.Auth(auth); err != nil {
		return err
	}

	// 设置发件人
	if err = client.Mail(from); err != nil {
		return err
	}

	// 添加收件人
	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}

	// 获取数据写入器
	w, err := client.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	// 写入邮件内容
	_, err = w.Write(msg)
	return err
}
