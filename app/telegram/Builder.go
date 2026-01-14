package telegram

import (
	"GO/app/core"
	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
	"GO/app/telegram/services"
)

type TelegramBuilder struct {
	Cntx *core.Context
}

func (f *TelegramBuilder) buildOutworldFacade() *outworld.OutworldFacade {
	return &outworld.OutworldFacade{Cntx: f.Cntx}
}

func (f *TelegramBuilder) buildUserRepository() *Repositories.UserRepository {
	repo := Repositories.UserRepository{Db: f.Cntx.GetDb()}
	return &repo
}

func (f *TelegramBuilder) buildReferalLinkRepository() *Repositories.ReferalLinkRepository {
	repo := Repositories.ReferalLinkRepository{Db: f.Cntx.GetDb()}
	return &repo
}

func (f *TelegramBuilder) buildReferalUserRepository() *Repositories.ReferalUserRepository {
	repo := Repositories.ReferalUserRepository{Db: f.Cntx.GetDb()}
	return &repo
}

func (f *TelegramBuilder) buildMessageService() *services.MessageService {
	repo := f.buildUserRepository()
	service := services.MessageService{UserRepo: repo, OutworldFacade: f.buildOutworldFacade(), CommandsHandler: f.buildCommandService()}
	return &service
}

func (f *TelegramBuilder) buildMyChatMemberService() *services.MyChatMemberService {
	return &services.MyChatMemberService{UserRepository: *f.buildUserRepository()}
}

func (f *TelegramBuilder) buildCommandService() *services.CommandService {
	return &services.CommandService{UserRepository: f.buildUserRepository(), OutworldFacade: f.buildOutworldFacade(), ReferalService: f.buildReferalService()}
}

func (f *TelegramBuilder) buildCallbackQueryService() *services.CallbackQueryService {
	return &services.CallbackQueryService{OutworldFacade: f.buildOutworldFacade()}
}

func (f *TelegramBuilder) buildReferalService() *services.ReferalService {
	return &services.ReferalService{UserRepository: f.buildUserRepository(), ReferalLinkRepository: f.buildReferalLinkRepository(), ReferalUserRepository: f.buildReferalUserRepository()}
}
