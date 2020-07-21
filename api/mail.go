package api

import (
	"fmt"
	"net/http"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sending mail")
}
