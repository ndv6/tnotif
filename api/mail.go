package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"time"

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

type templateData struct {
	Token string
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

		data := templateData{
			Token: req.Token,
		}

		subject := "Please verify your email"
		body, err := ParseTemplate("templates/template.html", data)
		if err != nil {
			fmt.Fprint(w, "cannot parse email template")
		}
		message := CreateEmailMessage(subject, body)

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
	logMail := models.LogMail{
		Email:  email,
		SentAt: time.Now(),
	}
	_, err := db.Exec("INSERT INTO log_mail(email, sent_at) VALUES ($1, $2)", logMail.Email, logMail.SentAt)
	return err
}

func ParseTemplate(filename string, data interface{}) (string, error) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), err
}

func CreateEmailMessage(subject string, body string) []byte {
	subject = "Subject: " + subject
	MIME := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + "\n" + MIME + "\n" + body)
	return message
}
