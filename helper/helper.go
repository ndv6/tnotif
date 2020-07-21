package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(varName string) string {
	godotenv.Load()
	return (os.Getenv(varName))
}
