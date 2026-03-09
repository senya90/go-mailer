package mailer

import (
	"fmt"
	"log/slog"
	"mail-service/internal/config"
	"mail-service/internal/models"
	"net/smtp"
	"strings"
)

type Mailer struct {
	cfg    *config.Config
	logger *slog.Logger
}

func NewMailer(cfg *config.Config, logger *slog.Logger) *Mailer {
	return &Mailer{
		cfg:    cfg,
		logger: logger,
	}
}

func (m *Mailer) SendEmail(params *models.SendEmailParams) error {
	m.logger.Info("Sending email", "to", params.To, "subject", params.Subject, "message", truncateMessage(params.Message, 10, 7))

	auth := smtp.PlainAuth("", m.cfg.SMTPFrom, m.cfg.SMTPPassword, m.cfg.SMTPHost)
	address := fmt.Sprintf("%s:%d", m.cfg.SMTPHost, m.cfg.SMTPPort)

	var mail strings.Builder

	fmt.Fprintf(&mail, "From: Mail service. no-reply <%s>\r\n", m.cfg.SMTPFrom)
	fmt.Fprintf(&mail, "To: %s\r\n", params.To)
	fmt.Fprintf(&mail, "Subject: %s\r\n", params.Subject)
	fmt.Fprintf(&mail, "MIME-version: 1.0;\r\n")
	fmt.Fprintf(&mail, "Content-Type: text/html; charset=\"UTF-8\";\r\n")
	fmt.Fprintf(&mail, "%s", params.Message)

	err := smtp.SendMail(address, auth, m.cfg.SMTPFrom, []string{params.To}, []byte(mail.String()))

	if err != nil {
		m.logger.Error("Failed to send email", "to", params.To, "error", err)
		return err
	}

	m.logger.Info("Email sent successfully", "to", params.To)
	return nil
}

func truncateMessage(message string, start int, end int) string {
	runes := []rune(message)
	if len(runes) <= start+end {
		return message
	}

	return string(runes[:start]) + "..." + string(runes[len(runes)-end:])
}
