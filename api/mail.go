package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/ndv6/tnotif/api/storage"

	"github.com/ndv6/tnotif/helper"
	"github.com/ndv6/tnotif/models"
)

type smtpServer struct {
	Host string
	Port string
}

type smtpEmail struct {
	Email    string
	Password string
}

type smtpRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (s *smtpServer) getAddress() string {
	return s.Host + ":" + s.Port
}

func SendMailHandler(db string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req smtpRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Fprint(w, fmt.Sprintf("%v", err))
		}

		sender := smtpEmail{
			Email:    helper.GetEnv("EMAIL_ACC"),
			Password: helper.GetEnv("EMAIL_PASSWORD"),
		}

		to := []string{
			req.Email,
		}

		server := smtpServer{
			Host: "smtp.gmail.com",
			Port: "587",
		}

		message := []byte(req.Token)

		auth := smtp.PlainAuth("", sender.Email, sender.Password, server.Host)
		err = smtp.SendMail(server.getAddress(), auth, sender.Email, to, message)
		if err != nil {
			fmt.Fprint(w, fmt.Sprintf("%v", err))
			return
		}
		database := storage.GetStorage(db)
		err = LogMail(req.Email, database)
		if err != nil {
			fmt.Fprint(w, "Cannot log the sent email")
			return
		}
		return
	})
}

func LogMail(email string, db storage.Storage) error {
	logMail := models.LogMail{
		Email:  email,
		SentAt: time.Now(),
	}
	err := db.Create(logMail)
	if err != nil {
		return err
	}
	return err
}
