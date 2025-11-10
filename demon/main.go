package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"GO/app/tel/updates/stru"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	http.HandleFunc("/webhook", httpHandler)
	go http.ListenAndServe(":8000", nil)
}

func httpHandler(resp http.ResponseWriter, req *http.Request) {
	var update stru.Update

	// Читаем Body правильно
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	// Декодируем JSON
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for _, value := range update.Result {
		spew.Dump(value)
		// fmt.Println(value.Message)
	}
}

func handleUpdates(updatesChannel chan stru.Update) {
	var UserChannels map[int]stru.UserChannel

	for update := range updatesChannel {
		checkI
	}
}
