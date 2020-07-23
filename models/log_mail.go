package models

import (
	"time"
)

type LogMail struct {
	Email  string    `json:"email"`
	SentAt time.Time `json:"sentAt"`
}

type SmtpResponse struct {
	Email string `json:"email"`
}

type SmtpRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
