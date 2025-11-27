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

	handleUpdates(value, &Context)
	resp.WriteHeader(http.StatusOK)
}

func handleUpdates(tg_update updates.Update, Context *core.Context) {
	user_id := tg_update.GetUserId()
	activeStruct, ok := UserChannels[user_id]
	if !ok {
		ch := make(chan *updates.Update)
		activeStruct = entities.UserChannel{Update: &tg_update, Ch: &ch}
		UserChannels[user_id] = activeStruct
		go handleUpdate(&ch, Context)
	}
	*activeStruct.Ch <- &tg_update

	// Здесь нужно будет разбить обновления где есть юзер а где нет
}

func handleUpdate(channel *chan *updates.Update, Context *core.Context) {
	facade := telegram.TelegramFacade{Cntx: Context}
	for item := range *channel {
		update_type := item.GetUpdateType()
		switch update_type {
		case updates.MessageType:
			facade.HandleMessageUpdate(*item.GetMessage())
		case updates.MyChatMemberType:
			facade.HandleMyChatMemberUpdate(*item.GetMyChatMember())
		}
	}
}
