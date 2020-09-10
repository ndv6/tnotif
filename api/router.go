package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/ndv6/tnotif/helper"
)

const (
	template_file = "templates/template.html"
)

func Router(db string) http.Handler {
	ss := SmtpService{
		SmtpSender: SMTPEmail{
			ApiKey: helper.GetEnv("PRIVATE_API_KEY"),
			Domain: helper.GetEnv("DOMAIN"),
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

	r.Group(func(r chi.Router) {
		r.Use(httprate.LimitByIP(1, 10*time.Second))
		r.Post("/sendMail", ss.SendMailHandler(db)) // Yuly Haruka
	})

	r.Get("/", Home)
	r.NotFound(NotFound)
	return r
}
