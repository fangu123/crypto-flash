package service

import (
	"fmt"
	"go-admin/common/util"

	ext "go-admin/config"
)

var emailClient *util.Email

// InitEmail 初始化邮件客户端
func InitEmail() {
	emailConfig := &util.EmailConfig{
		Host:     ext.ExtConfig.Email.Host,
		Port:     ext.ExtConfig.Email.Port,
		User:     ext.ExtConfig.Email.User,
		Password: ext.ExtConfig.Email.Password,
		From:     ext.ExtConfig.Email.From,
	}
	emailClient = util.NewEmail(emailConfig)
}

// SendEmail 发送邮件
// to: 收件人列表
// subject: 邮件主题
// body: 邮件内容
// isHTML: 是否为HTML格式
func SendEmail(to []string, subject, body string, isHTML bool) error {
	if emailClient == nil {
		return fmt.Errorf("邮件客户端未初始化")
	}
	return emailClient.Send(to, subject, body, isHTML)
}

// SendHTMLEmail 发送HTML格式邮件
func SendHTMLEmail(to []string, subject, htmlBody string) error {
	return SendEmail(to, subject, htmlBody, true)
}

// SendTextEmail 发送纯文本格式邮件
func SendTextEmail(to []string, subject, textBody string) error {
	return SendEmail(to, subject, textBody, false)
}
