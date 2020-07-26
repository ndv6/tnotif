package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ndv6/tnotif/api"
	"github.com/ndv6/tnotif/helper"
	"github.com/ndv6/tnotif/models"
)

func TestSendMail(t *testing.T) {
	bodyReq := models.SmtpRequest{
		Email: "testing@example.com",
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

	ss := api.SmtpService{
		Server: api.SmtpServer{
			Host: "smtp.gmail.com",
			Port: "587",
		},
		Template: helper.GetEnv("TEMPLATE"),
	}
	handler := http.HandlerFunc(ss.SendMailHandler("mock"))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected HTTP Status Code 200, got: %v", rr.Code)
	}

	res, err := json.Marshal(models.SmtpResponse{
		Email: bodyReq.Email,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bytes.TrimRight(rr.Body.Bytes(), "\n"), res) {
		t.Fatalf("Expected email %v, got: %v", string(res), string(rr.Body.Bytes()))
	}
}
