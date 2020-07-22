package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ndv6/tnotif/api"
)

func TestSendMail(t *testing.T) {
	bodyReq := api.SmtpRequest{
		Email: "yuly.ocbcnisp@gmail.com",
		Token: "testing",
	}
	reqJSON, err := json.Marshal(bodyReq)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/sendMail", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.SendMailHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected HTTP Status Code 200, got: %v", rr.Code)
	}

	res, err := json.Marshal(api.SmtpResponse{
		Email: bodyReq.Email,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bytes.TrimRight(rr.Body.Bytes(), "\n"), res) {
		t.Fatalf("Expected email %v, got: %v", res, rr.Body.Bytes())
	}
}
