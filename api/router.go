package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

const (
	template_file = "templates/template.html"
)

type SmtpServer struct {
	Host string
	Port string
}

func (s *SmtpServer) getAddress() string {
	return s.Host + ":" + s.Port
}

func Router(db string) http.Handler {
	ss := SmtpService{
		Server: SmtpServer{
			Host: "smtp.gmail.com",
			Port: "587",
		},
		Template: template_file,
	}

	r := chi.NewRouter()

	// Yuly Haruka
	r.Post("/sendMail", ss.SendMailHandler(db))
	r.NotFound(NotFound)
	return r
}
