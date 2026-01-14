package telegram

import (
	"GO/app/telegram/entities"
	"GO/app/telegram/updates"
)

type TelegramFacade struct {
	Builder *TelegramBuilder
}

func (f *TelegramFacade) HandleMessageUpdate(update updates.Message, jobsChannel chan entities.Job) {
	service := f.Builder.buildMessageService()
	service.HandleMessageUpdate(update, jobsChannel)
}

func (f *TelegramFacade) HandleMyChatMemberUpdate(update updates.MyChatMember, jobsChannel chan entities.Job) {
	service := f.Builder.buildMyChatMemberService()
	service.HandleMyChatMemberUpdate(update, jobsChannel)
}

func (f *TelegramFacade) HandleCallbackQuery(update updates.CallbackQuery, jobsChannel chan entities.Job) {
	service := f.Builder.buildCallbackQueryService()
	service.HandleCallbackQuery(update, jobsChannel)
}
