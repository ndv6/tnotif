package helper

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr     string `json:"addr"`
	Database string `json:"database"`
}

func GetEnv(varName string) string {
	godotenv.Load()
	return (os.Getenv(varName))
}

func LoadConfig(file string) (Config, error) {
	var cfg Config
	f, err := os.Open(file)
	if err != nil {
		return Config{}, err
	}
	err = json.NewDecoder(f).Decode(&cfg)
	return cfg, err
}
