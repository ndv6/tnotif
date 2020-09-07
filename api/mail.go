package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/ndv6/tnotif/constants"

	"github.com/ndv6/tnotif/api/storage"
	"github.com/ndv6/tnotif/helper"
	"github.com/ndv6/tnotif/models"
)

type smtpEmail struct {
	ApiKey string
	Domain string
}

type templateData struct {
	Token string
}

type SmtpService struct {
	SmtpEmail smtpEmail
	Template  string
}

func (ss *SmtpService) SendMailHandler(db string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req models.SmtpRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.Header().Set(constants.ContentType, constants.JSON)
			helper.HTTPError(w, http.StatusBadRequest, "Cannot parse request")
			return
		}

		mailgun := smtpEmail{
			ApiKey: helper.GetEnv("PRIVATE_API_KEY"),
			Domain: helper.GetEnv("DOMAIN"),
		}

		body, err := ParseTemplate("templates/template.html", templateData{Token: req.Token})
		if err != nil {
			w.Header().Set(constants.ContentType, constants.JSON)
			helper.HTTPError(w, http.StatusInternalServerError, constants.FailedParseTemplate)
			return
		}
		// message := CreateEmailMessage(subject, body)

		status, err := helper.SendMessage(mailgun.ApiKey, mailgun.Domain, helper.GetEnv("EMAIL_ACC"), req.Email, constants.SubjectEmail, body)
		if err != nil {
			helper.SendMessageToTelegram(r, http.StatusInternalServerError, constants.SendMailFailed)
			w.Header().Set(constants.ContentType, constants.JSON)
			helper.HTTPError(w, http.StatusInternalServerError, constants.SendMailFailed)
			return
		}

		database := storage.GetStorage(db)
		err = LogMail(req.Email, database)
		if err != nil {
			helper.SendMessageToTelegram(r, http.StatusInternalServerError, constants.FailedConnectDatabase)
			w.Header().Set(constants.ContentType, constants.JSON)
			helper.HTTPError(w, http.StatusInternalServerError, constants.FailedConnectDatabase)
			return
		}

		objResponse := models.SmtpResponse{
			Email: req.Email,
		}
		_, res, err := helper.NewResponseBuilder(w, true, constants.SendMailSuccess, objResponse)
		if err != nil {
			w.Header().Set(constants.ContentType, constants.JSON)
			helper.HTTPError(w, http.StatusBadRequest, constants.CannotEncodeResponse)
			return
		}

		w.Header().Set(constants.ContentType, constants.JSON)
		fmt.Fprintf(w, res)
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
