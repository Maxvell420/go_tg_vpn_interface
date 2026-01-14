package telegram

import (
	"GO/app/core"
	"GO/app/domain/User"
	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
	"GO/app/telegram/services"
)

type TelegramBuilder struct {
	Cntx *core.Context
}

func (f *TelegramBuilder) BuildOutworldFacade() *outworld.OutworldFacade {
	return &outworld.OutworldFacade{Builder: &outworld.Builder{Cntx: f.Cntx}}
}

func (f *TelegramBuilder) buildUserRepository() *Repositories.UserRepository {
	repo := Repositories.UserRepository{Db: f.Cntx.GetDb()}
	return &repo
}

func (f *TelegramBuilder) buildMessageService() *services.MessageService {
	service := services.MessageService{UserFacade: f.buildUserFacade(), OutworldFacade: f.BuildOutworldFacade(), CommandsHandler: f.buildCommandService()}
	return &service
}

func (f *TelegramBuilder) buildMyChatMemberService() *services.MyChatMemberService {
	return &services.MyChatMemberService{UserRepository: *f.buildUserRepository()}
}

func (f *TelegramBuilder) buildCommandService() *services.CommandService {
	return &services.CommandService{UserFacade: f.buildUserFacade(), OutworldFacade: f.BuildOutworldFacade()}
}

func (f *TelegramBuilder) buildCallbackQueryService() *services.CallbackQueryService {
	return &services.CallbackQueryService{OutworldFacade: f.BuildOutworldFacade()}
}

func (f *TelegramBuilder) buildUserFacade() *user.UserFacade {
	return &user.UserFacade{Builder: &user.UserBuilder{Cntx: f.Cntx}}
}
