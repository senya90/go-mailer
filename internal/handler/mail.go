package handler

import (
	"encoding/json"
	"log"
	"mail-service/internal/mailer"
	"mail-service/internal/models"
	"net/http"
)

type MailHandler struct {
	mailer *mailer.Mailer
}

func NewMailHandler(m *mailer.Mailer) *MailHandler {
	return &MailHandler{
		mailer: m,
	}
}

func (handler *MailHandler) Send(w http.ResponseWriter, r *http.Request) {
	var params models.SendEmailParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// todo: добавить библиотеку для валидации
	if params.To == "" || params.Message == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if params.Subject == "" {
		params.Subject = "Mail service notification"
	}

	sendErr := handler.mailer.SendEmail(&params)

	if sendErr != nil {
		log.Printf("Failed to send email to %s: %v", params.To, sendErr)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)

		return
	}

	log.Printf("Email sent to %s", params.To)
	w.WriteHeader(http.StatusOK)
}
