package handler

import (
	"encoding/json"
	"log"
	"mail-service/internal/mailer"
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

type sendEmailParams struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

func (handler *MailHandler) Send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var params sendEmailParams
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

	sendErr := handler.mailer.SendEmail(params.To, params.Message)

	if sendErr != nil {
		log.Printf("Failed to send email to %s: %v", params.To, sendErr)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)

		return
	}

	log.Printf("Email sent to %s", params.To)
	w.WriteHeader(http.StatusOK)
}
