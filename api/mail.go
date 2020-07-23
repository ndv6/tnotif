package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
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

type SmtpRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type templateData struct {
	Token string
}

type SmtpResponse struct {
	Email string `json:"email"`
}

func (s *smtpServer) getAddress() string {
	return s.Host + ":" + s.Port
}

func SendMailHandler(db string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req SmtpRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			helper.HTTPError(w, http.StatusBadRequest, "Cannot parse request")
			return
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
			helper.HTTPError(w, http.StatusBadRequest, "Cannot parse email template")
			return
		}
		message := CreateEmailMessage(subject, body)

		auth := smtp.PlainAuth("", sender.Email, sender.Password, server.Host)
		err = smtp.SendMail(server.getAddress(), auth, sender.Email, to, message)
		if err != nil {
			helper.HTTPError(w, http.StatusBadRequest, "Failed to send mail")
			return
		}
		database := storage.GetStorage(db)
		err = LogMail(req.Email, database)
		if err != nil {
			fmt.Fprint(w, "Cannot log the sent email")
			return
		}
		resp := SmtpResponse{
			Email: req.Email,
		}

		for _, e := range to {
			err = LogMail(e, database)
			if err != nil {
				helper.HTTPError(w, http.StatusBadRequest, "Can not log send mail")
				return
			}
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			helper.HTTPError(w, http.StatusBadRequest, "Can not encode response")
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
