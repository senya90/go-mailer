package mailer

import (
	"fmt"
	"mail-service/internal/config"
	"mail-service/internal/models"
	"net/smtp"
	"strings"
)

type Mailer struct {
	cfg *config.Config
}

func NewMailer(cfg *config.Config) *Mailer {
	return &Mailer{
		cfg: cfg,
	}
}

func (m *Mailer) SendEmail(params *models.SendEmailParams) error {
	auth := smtp.PlainAuth("", m.cfg.SMTPFrom, m.cfg.SMTPPassword, m.cfg.SMTPHost)
	address := fmt.Sprintf("%s:%d", m.cfg.SMTPHost, m.cfg.SMTPPort)

	var mail strings.Builder

	fmt.Fprintf(&mail, "From: Mail service. no-reply <%s>\r\n", m.cfg.SMTPFrom)
	fmt.Fprintf(&mail, "To: %s\r\n", params.To)
	fmt.Fprintf(&mail, "Subject: %s\r\n", params.Subject)
	fmt.Fprintf(&mail, "MIME-version: 1.0;\r\n")
	fmt.Fprintf(&mail, "Content-Type: text/html; charset=\"UTF-8\";\r\n")
	fmt.Fprintf(&mail, "%s", params.Message)

	return smtp.SendMail(address, auth, m.cfg.SMTPFrom, []string{params.To}, []byte(mail.String()))
}
