package api

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Print("this is for sending mail")
	return
}
