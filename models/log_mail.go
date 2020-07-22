package models

import (
	"time"
)

type LogMail struct {
	Email  string    `json:"email"`
	SentAt time.Time `json:"sentAt"`
}
