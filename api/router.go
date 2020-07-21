package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/sendMail", SendMailHandler)
	r.NotFound(NotFound)
	return r
}
