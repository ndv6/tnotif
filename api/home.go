package api

import (
	"net/http"

	"github.com/ndv6/tnotif/helper"
)

func Home(w http.ResponseWriter, r *http.Request) {
	helper.HTTPError(w, http.StatusOK, "Welcome to Tnotif")
	return
}
