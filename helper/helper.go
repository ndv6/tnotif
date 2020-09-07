package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	mailgun "github.com/mailgun/mailgun-go/v4"
)

type Config struct {
	Addr     string `json:"addr"`
	Database string `json:"database"`
}

func GetEnv(varName string) string {
	godotenv.Load()
	return (os.Getenv(varName))
}

func HTTPError(w http.ResponseWriter, status int, errorMessage string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": errorMessage})
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

func SendMessageToTelegram(r *http.Request, status int, errorMessage string) error {

	current_time := time.Now()
	chat_id := os.Getenv("CHATID")
	text := "There has been an exception.\n" +
		"<b>HTTP Status</b>:" + strconv.Itoa(status) + "\n" +
		"<b>Message</b> : " + errorMessage + "\n" +
		"<b>Timestamp</b> :" + current_time.Format(time.RFC1123) + "\n" +
		"<b>Endpoint</b> :" + html.EscapeString(r.URL.Path) + "\n" +
		"<b>Method</b> :" + r.Method
	data, err := json.Marshal(map[string]string{
		"chat_id":    chat_id,
		"text":       text,
		"parse_mode": "HTML",
	})
	if err != nil {
		return err
	}

	url_bot_telegram := os.Getenv("TELEGRAM")
	resp, err := http.Post(url_bot_telegram, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}

func SendMessage(apiKey, domain, sender, recipient, subject, body string) error {
	mg := mailgun.NewMailgun(domain, apiKey)
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
