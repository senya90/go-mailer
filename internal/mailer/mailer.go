package mailer

import (
	"fmt"
	"mail-service/internal/config"
	"net/smtp"
)

type Mailer struct {
	cfg *config.Config
}

func NewMailer(cfg *config.Config) *Mailer {
	return &Mailer{
		cfg: cfg,
	}
}

func (m *Mailer) SendEmail(to string, message string) error {
	auth := smtp.PlainAuth("", m.cfg.SMTPFrom, m.cfg.SMTPPassword, m.cfg.SMTPHost)
	address := fmt.Sprintf("%s:%d", m.cfg.SMTPHost, m.cfg.SMTPPort)

	mail := []byte(
		"From: " + m.cfg.SMTPFrom + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: Email sender\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			"\r\n" +
			message,
	)

	return smtp.SendMail(address, auth, m.cfg.SMTPFrom, []string{to}, []byte(mail))
}
