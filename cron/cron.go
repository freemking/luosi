package cron

import (
	"fastener-pro/config"
	"fastener-pro/mail"
	"fastener-pro/models"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

// SetupCron 配置定时任务
func SetupCron(cfg config.Config) {
	c := cron.New()

	// 每天发送一次联系表单数据到邮箱
	_, err := c.AddFunc("0 0 * * *", func() {
		SendContactEmails(cfg)
	})
	if err != nil {
		log.Printf("Failed to add cron job: %v", err)
		return
	}

	// 启动定时任务
	c.Start()
	log.Println("Cron jobs started")
}

// SendContactEmails 发送联系表单数据到邮箱
func SendContactEmails(cfg config.Config) {
	// 获取所有联系表单数据
	contacts, err := models.GetContacts()
	if err != nil {
		log.Printf("Failed to get contacts: %v", err)
		return
	}

	if len(contacts) == 0 {
		log.Println("No contact forms to send")
		return
	}

	// 构建邮件内容
	var body strings.Builder
	body.WriteString("Contact Form Submissions\n\n")
	body.WriteString(fmt.Sprintf("Total submissions: %d\n\n", len(contacts)))

	for i, contact := range contacts {
		body.WriteString(fmt.Sprintf("Submission #%d\n", i+1))
		body.WriteString(fmt.Sprintf("Name: %s\n", contact.Name))
		body.WriteString(fmt.Sprintf("Email: %s\n", contact.Email))
		body.WriteString(fmt.Sprintf("Phone: %s\n", contact.Phone))
		body.WriteString(fmt.Sprintf("Company: %s\n", contact.Company))
		body.WriteString(fmt.Sprintf("Product: %s\n", contact.Product))
		body.WriteString(fmt.Sprintf("Message: %s\n", contact.Message))
		body.WriteString(fmt.Sprintf("Submitted at: %s\n\n", contact.CreatedAt))
	}

	// 发送邮件
	mailConfig := mail.MailConfig{
		SMTPHost:     cfg.Mail.SMTPHost,
		SMTPPort:     cfg.Mail.SMTPPort,
		SMTPUsername: cfg.Mail.SMTPUsername,
		SMTPPassword: cfg.Mail.SMTPPassword,
		From:         cfg.Mail.From,
		To:           cfg.Mail.To,
	}

	subject := fmt.Sprintf("Contact Form Submissions - %s", time.Now().Format("2006-01-02"))
	err = mail.SendEmail(mailConfig, subject, body.String())
	if err != nil {
		log.Printf("Failed to send contact emails: %v", err)
		return
	}

	log.Printf("Successfully sent %d contact form submissions via email", len(contacts))
}
