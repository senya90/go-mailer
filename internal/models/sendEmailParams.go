package models

type SendEmailParams struct {
	To      string `json:"to" validate:"required,email"`
	Message string `json:"message" validate:"required,max=5000"`
	Subject string `json:"subject" validate:"omitempty,min=1,max=500"`
}
