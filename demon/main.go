package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"GO/app/tel/updates/interfaces"
	"GO/app/tel/updates/stru"
	// "GO/app/tel/updates/stru/up_types"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

var UserChannels map[int]stru.UserChannel

func main() {
	godotenv.Load()
	http.HandleFunc("/webhook", httpHandler)
	http.ListenAndServe(":8000", nil)
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
		handleUpdates(value)
		// fmt.Println(value.Message)
	}
}

func handleUpdates(update stru.Result) {
	value := update.Message

	if value != nil {
		handleMessage(value, UserChannels)
	}
}

func handleMessage(update interfaces.UserGetter, channels map[int]stru.UserChannel) {
	var activeStruct stru.UserChannel

	user_id = update.GetUser()

	activeStruct, ok := channels[user_id]
}
