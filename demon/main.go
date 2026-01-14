package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"GO/app/core"
	"GO/app/telegram"
	"GO/app/telegram/entities"
	"GO/app/telegram/updates"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

var (
	UserChannels map[int]entities.UserChannel
	Context      core.Context
	jobsChannel  chan entities.Job
)

func main() {
	godotenv.Load("../.env")
	http.HandleFunc("/webhook", httpHandler)
	UserChannels = make(map[int]entities.UserChannel)
	jobsChannel = make(chan entities.Job, 10)
	go createCronJobs(jobsChannel)
	go handleJobs(jobsChannel)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		spew.Dump(err)
	}
}

func createCronJobs(jobsChannel chan entities.Job) {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		jobsChannel <- entities.Job{Type: entities.TrafficUsage}
	}
}

func handleJob(job entities.Job) {
	switch job.Type {
	case entities.TrafficUsage:
		go handleTrafficUsage(job)
	}
}

func handleJobs(jobsChannel chan entities.Job) {
	for job := range jobsChannel {
		handleJob(job)
	}
}

func handleTrafficUsage(job entities.Job) {
	// builder := telegram.TelegramBuilder{Cntx: &Context}
	// facade := telegram.TelegramFacade{Builder: &builder}
	// facade.HandleTrafficUsage(job)
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

	handleUpdates(value, &Context, jobsChannel)
	resp.WriteHeader(http.StatusOK)
}

func handleUpdates(tg_update updates.Update, Context *core.Context, jobsChannel chan entities.Job) {
	user_id := tg_update.GetUserId()
	activeStruct, ok := UserChannels[user_id]
	if !ok {
		ch := make(chan *updates.Update)
		activeStruct = entities.UserChannel{Update: &tg_update, Ch: &ch}
		UserChannels[user_id] = activeStruct
		go handleUpdate(&ch, Context, jobsChannel)
	}
	*activeStruct.Ch <- &tg_update

	// Здесь нужно будет разбить обновления где есть юзер а где нет
}

func handleUpdate(channel *chan *updates.Update, Context *core.Context, jobsChannel chan entities.Job) {
	builder := telegram.TelegramBuilder{Cntx: Context}
	facade := telegram.TelegramFacade{Builder: &builder}
	for item := range *channel {
		update_type := item.GetUpdateType()
		switch update_type {
		case updates.MessageType:
			facade.HandleMessageUpdate(*item.GetMessage(), jobsChannel)
		case updates.MyChatMemberType:
			facade.HandleMyChatMemberUpdate(*item.GetMyChatMember(), jobsChannel)
		case updates.CallbackQueryType:
			facade.HandleCallbackQuery(*item.GetCallbackQuery(), jobsChannel)
		}
	}
}
