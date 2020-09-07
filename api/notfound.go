package api

import (
	"net/http"

	"github.com/ndv6/tnotif/helper"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	helper.HTTPError(w, http.StatusNotFound, "Not Found")
	return
}
