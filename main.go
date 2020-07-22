package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ndv6/tnotif/helper"

	_ "github.com/lib/pq"
	"github.com/ndv6/tnotif/api"
)

func main() {
	cfg, err := helper.LoadConfig("config.json")
	if err != nil {
		log.Fatal("unable to load configuration config.json")
	}

	db := helper.GetEnv("DB")
	if db == "" {
		db = "mock"
	}
	fmt.Println("Serving at port :8082")
	err = http.ListenAndServe(cfg.Addr, api.Router(db))
	if err != nil {
		log.Fatal(err)
	}
}
