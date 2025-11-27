package telegram

import (
	"GO/app/core"
	"GO/app/outworld"
	"GO/app/telegram/repositories"
	"GO/app/telegram/services"
	"GO/app/telegram/updates"
)

type TelegramFacade struct {
	Cntx *core.Context
}

func (f *TelegramFacade) HandleMessageUpdate(update updates.Message) {
	service := f.buildMessageService()
	service.HandleMessageUpdate(update)
}

func (f *TelegramFacade) HandleMyChatMemberUpdate(update updates.MyChatMember) {
	service := f.buildMyChatMemberService()
	service.HandleMyChatMemberUpdate(update)
}

func (f *TelegramFacade) buildOutworldFacade() *outworld.OutworldFacade {
	return &outworld.OutworldFacade{Cntx: f.Cntx}
}

func (f *TelegramFacade) buildUserRepository() *repositories.UserRepository {
	repo := repositories.UserRepository{Db: f.Cntx.GetDb()}
	return &repo
}

func (f *TelegramFacade) buildMessageService() *services.MessageService {
	repo := f.buildUserRepository()
	service := services.MessageService{UserRepo: repo, OutworldFacade: f.buildOutworldFacade()}
	return &service
}

func (f *TelegramFacade) buildMyChatMemberService() *services.MyChatMemberService {
	return &services.MyChatMemberService{UserRepository: *f.buildUserRepository()}
}
