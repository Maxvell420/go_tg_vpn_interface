package telegram

import (
	"GO/app/telegram/entities"
)

type TelegramCronFacade struct {
	Builder *TelegramBuilder
}

func (f *TelegramCronFacade) HandleTrafficUsage(job entities.Job) {
	service := f.Builder.buildTrafficService()
	service.HandleTrafficUsage(job)
}
