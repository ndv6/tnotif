package api

import (
	"net/http"

	"github.com/ndv6/tnotif/constants"
	"github.com/ndv6/tnotif/helper"
)

func Home(w http.ResponseWriter, r *http.Request) {
	helper.NewResponseBuilder(w, true, constants.Welcome, nil)
	return
}
