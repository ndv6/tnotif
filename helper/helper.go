package helper

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(varName string) string {
	godotenv.Load()
	return (os.Getenv(varName))
}

func HTTPError(w http.ResponseWriter, status int, errorMessage string) {
	json.NewEncoder(w).Encode(map[string]string{"error": errorMessage})
}
