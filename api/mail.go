package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/ndv6/tnotif/helper"
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

func SendMailHandler(db *sql.DB) http.HandlerFunc {
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

		for _, e := range to {
			LogMail(e, db)
		}
		return
	})
}

func LogMail(email string, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO log_mail(email, sent_at) VALUES ($1, $2)", email, time.Now())
	return err
}
