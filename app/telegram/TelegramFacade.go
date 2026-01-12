package telegram

import (
	"GO/app/core"
	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
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

func (f *TelegramFacade) HandleCallbackQuery(update updates.CallbackQuery) {
	service := f.buildCallbackQueryService()
	service.HandleCallbackQuery(update)
}

func (f *TelegramFacade) buildOutworldFacade() *outworld.OutworldFacade {
	return &outworld.OutworldFacade{Cntx: f.Cntx}
}

func (f *TelegramFacade) buildUserRepository() *Repositories.UserRepository {
	repo := Repositories.UserRepository{Db: f.Cntx.GetDb()}
	return &repo
}

func (f *TelegramFacade) buildReferalLinkRepository() *Repositories.ReferalLinkRepository {
	repo := Repositories.ReferalLinkRepository{Db: f.Cntx.GetDb()}
	return &repo
}

func (f *TelegramFacade) buildReferalUserRepository() *Repositories.ReferalUserRepository {
	repo := Repositories.ReferalUserRepository{Db: f.Cntx.GetDb()}
	return &repo
}

func (f *TelegramFacade) buildMessageService() *services.MessageService {
	repo := f.buildUserRepository()
	service := services.MessageService{UserRepo: repo, OutworldFacade: f.buildOutworldFacade(), CommandsHandler: f.buildCommandService()}
	return &service
}

func (f *TelegramFacade) buildMyChatMemberService() *services.MyChatMemberService {
	return &services.MyChatMemberService{UserRepository: *f.buildUserRepository()}
}

func (f *TelegramFacade) buildCommandService() *services.CommandService {
	return &services.CommandService{UserRepository: f.buildUserRepository(), OutworldFacade: f.buildOutworldFacade(), ReferalLinkRepository: f.buildReferalLinkRepository(), ReferalUserRepository: f.buildReferalUserRepository()}
}

func (f *TelegramFacade) buildCallbackQueryService() *services.CallbackQueryService {
	return &services.CallbackQueryService{OutworldFacade: f.buildOutworldFacade()}
}
