package telegram

import (
	"bytes"
	"encoding/json"

	// "fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type Request struct {
	BotToken *string
}

type GetRequest struct {
	method  TelegramRequest
	offset  int
	timeout int
}

type PostRequest struct {
	Method  TelegramRequest
	Message *Message
}

func (r *Request) SendGet(data GetRequest) string {
	url := "https://api.telegram.org/bot" + *r.BotToken + "/" + string(data.method)
	resp, err := http.Get(url)
	if err != nil {
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func (r *Request) SendPost(data PostRequest) string {
	url := "https://api.telegram.org/bot" + *r.BotToken + "/" + string(data.Method)
	jsonData, err := json.Marshal(data.Message)
	spew.Dump(string(jsonData))
	if err != nil {
		// обработать ошибки
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	spew.Dump(string(body))
	return string(body)
}

type TelegramRequest string

const (
	GetUpdates  TelegramRequest = "GetUpdates"
	SendMessage TelegramRequest = "sendMessage"
)
