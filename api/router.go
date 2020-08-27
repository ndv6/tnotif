package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
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
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
	}))

	r.Use(httprate.LimitByIP(1, 10*time.Second))

	r.Post("/sendMail", ss.SendMailHandler(db)) // Yuly Haruka
	r.Get("/", Home)
	r.NotFound(NotFound)
	return r
}
