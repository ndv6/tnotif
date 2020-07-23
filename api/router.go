package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router(db string) http.Handler {
	r := chi.NewRouter()
	r.Post("/sendMail", SendMailHandler(db))
	r.NotFound(NotFound)
	return r
}
