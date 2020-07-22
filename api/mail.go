package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"

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

type SmtpRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type SmtpResponse struct {
	Email string `json:"email"`
}

func (s *smtpServer) getAddress() string {
	return s.Host + ":" + s.Port
}

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	var req SmtpRequest
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

	resp := SmtpResponse{
		Email: req.Email,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Fprint(w, "Unable to encode response")
		return
	}
	return
}
