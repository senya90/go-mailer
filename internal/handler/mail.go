package handler

import (
	"encoding/json"
	"log/slog"
	"mail-service/internal/mailer"
	"mail-service/internal/models"
	"mail-service/internal/validator"
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
	apiKey := r.Header.Get("X-Email-Api-Key")

	if apiKey != handler.mailer.Cfg.ApiKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var params models.SendEmailParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		handler.logger.Error("Invalid request body", "params", params)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if params.Subject == "" {
		params.Subject = "Mail service notification"
	}

	err = validator.Validate(params)
	if err != nil {
		handler.logger.Error("Validation failed", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.mailer.SendEmail(&params)
	if err != nil {
		handler.logger.Error("Failed to send email", "to", params.To, "error", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
