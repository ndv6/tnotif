package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Post("/sendMail", SendMailHandler(db))
	r.NotFound(NotFound)
	return r
}
