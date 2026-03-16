package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

// MailConfig 邮件配置
type MailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	From         string
	To           string
}

// SendEmail 发送邮件
func SendEmail(config MailConfig, subject, body string) error {
	// 构建邮件内容
	message := fmt.Sprintf("From: %s\r\n", config.From)
	message += fmt.Sprintf("To: %s\r\n", config.To)
	message += fmt.Sprintf("Subject: %s\r\n\r\n", subject)
	message += body

	// 设置认证信息
	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost)

	// 构建SMTP服务器地址
	smtpAddr := fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort)

	// 发送邮件
	err := smtp.SendMail(smtpAddr, auth, config.From, strings.Split(config.To, ","), []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// FormatContactEmail 格式化联系表单邮件
func FormatContactEmail(name, email, phone, company, product, message string) string {
	body := fmt.Sprintf("New Contact Form Submission\n\n")
	body += fmt.Sprintf("Name: %s\n", name)
	body += fmt.Sprintf("Email: %s\n", email)
	body += fmt.Sprintf("Phone: %s\n", phone)
	body += fmt.Sprintf("Company: %s\n", company)
	body += fmt.Sprintf("Product Interest: %s\n", product)
	body += fmt.Sprintf("Message:\n%s\n", message)

	return body
}
