package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ndv6/tnotif/api"
)

type Config struct {
	Addr string `json:"addr"`
}

func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("unable to load configuration config.json")
	}

	fmt.Println("Serving at port :8082")
	err = http.ListenAndServe(cfg.Addr, api.Router())
	if err != nil {
		log.Fatal(err)
	}
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
