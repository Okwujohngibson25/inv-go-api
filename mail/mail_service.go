package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strconv"

	"github.com/coinserveringo/config"
	"gopkg.in/gomail.v2"
)

type MailService struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewMailService(cfg *config.Config) *MailService {
	port, _ := strconv.Atoi(cfg.SMTPPort)
	return &MailService{
		host:     cfg.SMTPHost,
		port:     port,
		username: cfg.SMTPUser,
		password: cfg.SMTPPass,
		from:     cfg.SMTPFrom,
	}
}

func (m *MailService) SendMail(to, subject, templateName string, data any) error {
	// Load and parse the HTML template
	tmplPath := filepath.Join("mail", "templates", templateName)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("error loading template: %w", err)
	}

	// Execute the template into a buffer
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	// Build the email
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body.String())

	// Send the mail
	dialer := gomail.NewDialer(m.host, m.port, m.username, m.password)
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
