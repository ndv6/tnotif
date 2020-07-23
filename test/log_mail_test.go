package test

import (
	"testing"

	"github.com/ndv6/tnotif/api"
	"github.com/ndv6/tnotif/api/storage"
)

func LogMailTest(t *testing.T) {
	database := storage.GetStorage("mock")
	email := "testing@example.com"

	err := api.LogMail(email, database)
	if err != nil {
		t.Fatal(err)
	}

	listLog, err := database.List()
	if len(listLog) != 1 {
		t.Fatalf("Expected 1 data log, got: %v", len(listLog))
	}
}
