package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	// "github.com/davecgh/go-spew/spew"

	"GO/app/core"
	"GO/app/telegram"
	"GO/app/telegram/entities"
	"GO/app/telegram/updates"

	"github.com/joho/godotenv"
)

var (
	UserChannels map[int]entities.UserChannel
	Context      core.Context
)

func main() {
	godotenv.Load("../.env")
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

	handleUpdates(&value, &Context)
	resp.WriteHeader(http.StatusOK)
}

func handleUpdates(tg_update updates.UserUpdate, Context *core.Context) {
	update_type := tg_update.GetUpdateType()

	update := tg_update.(*updates.Update)
	switch update_type {
	case updates.MessageType:

		user_id := update.Message.GetUser()
		activeStruct, ok := UserChannels[user_id]
		if !ok {
			ch := make(chan *updates.Update)
			activeStruct = entities.UserChannel{Update: update, Ch: ch}
			UserChannels[user_id] = activeStruct
			go handleUpdate(ch, Context)
			activeStruct.Ch <- update

		} else {
			activeStruct.Ch <- update
		}
	}
}

func handleUpdate(channel chan *updates.Update, Context *core.Context) {
	db := Context.GetDb()
	for item := range channel {
		facade := telegram.TelegramFacade{Db: db}
		if item.GetUpdateType() == updates.MessageType {
			facade.HandleMessageUpdate(*item.Message)
		}
	}
}
