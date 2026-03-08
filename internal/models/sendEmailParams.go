package models

type SendEmailParams struct {
	To      string `json:"to"`
	Message string `json:"message"`
	Subject string `json:"subject,omitempty"`
}
