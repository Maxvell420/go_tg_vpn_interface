package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	// "GO/app/tel/updates/stru/up_types"

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

	value := update.Result

	handleUpdates(&value)
}

func handleUpdates(update interfaces.Update) {
	update_type := update.GetUpdateType()

	switch update_type {
	case interfaces.Message:

		update := update.(*up_types.Message)
		handleMessage(update, UserChannels)
	}
}

func handleMessage(update up_types.Message, channels map[int]stru.UserChannel) {
	user_id := update.GetUser()

	activeStruct, ok := channels[user_id]

	if !ok {
		activeStruct := stru.UserChannel{Update: update}
	}
}
