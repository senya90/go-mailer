package handler

import (
	"encoding/json"
	"log/slog"
	"mail-service/internal/mailer"
	"mail-service/internal/models"
	"net/http"
)

type MailHandler struct {
	mailer *mailer.Mailer
	logger *slog.Logger
}

func NewMailHandler(m *mailer.Mailer, logger *slog.Logger) *MailHandler {
	return &MailHandler{
		mailer: m,
		logger: logger,
	}
}

func (handler *MailHandler) Send(w http.ResponseWriter, r *http.Request) {
	var params models.SendEmailParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		handler.logger.Error("Invalid request body", "params", params)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// todo: добавить библиотеку для валидации
	if params.To == "" || params.Message == "" {
		handler.logger.Error("Missing required fields", "params", params)
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if params.Subject == "" {
		params.Subject = "Mail service notification"
	}

	sendErr := handler.mailer.SendEmail(&params)

	if sendErr != nil {
		handler.logger.Error("Failed to send email", "to", params.To, "error", sendErr)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
