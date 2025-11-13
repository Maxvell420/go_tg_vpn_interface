package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"GO/app/telegram/entities"
	"GO/app/telegram/updates"

	"github.com/joho/godotenv"
)

var UserChannels map[int]entities.UserChannel

func main() {
	godotenv.Load()
	http.HandleFunc("/webhook", httpHandler)
	UserChannels = make(map[int]entities.UserChannel)
	http.ListenAndServe(":8000", nil)
}

func httpHandler(resp http.ResponseWriter, req *http.Request) {
	var update updates.TelegramUpdate

	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	err = json.Unmarshal(bytes, &update)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	value := update.Result

	handleUpdates(&value)
}

func handleUpdates(tg_update updates.UserUpdate) {
	update_type := tg_update.GetUpdateType()

	switch update_type {
	case updates.MessageType:

		update := tg_update.(*updates.Update)
		user_id := update.Message.GetUser()
		activeStruct, ok := UserChannels[user_id]
		if !ok {
			ch := make(chan updates.UserUpdate)
			activeStruct = entities.UserChannel{Update: update, Ch: &ch}
			UserChannels[user_id] = activeStruct
			go handleUpdate(&ch)
			fmt.Println("created")
			*activeStruct.Ch <- tg_update

		} else {
			fmt.Println("existed")
			*activeStruct.Ch <- tg_update
		}
	}
}

func handleUpdate(channel *chan updates.UserUpdate) {
	for item := range *channel {
		spew.Dump(item)
	}
}
