package telegram

import (
	"GO/app/core/database"
	"GO/app/telegram/repositories"
	"GO/app/telegram/services"
	"GO/app/telegram/updates"
)

type TelegramFacade struct {
	Db *database.Mysql
}

func (f *TelegramFacade) HandleMessageUpdate(update updates.Message) {
	service := f.buildMessageService()
	service.HandleMessageUpdate(update)
}

func (f *TelegramFacade) buildUserRepository() *repositories.UserRepository {
	repo := repositories.UserRepository{Db: f.Db}
	return &repo
}

func (f *TelegramFacade) buildMessageService() *services.MessageService {
	repo := f.buildUserRepository()
	service := services.MessageService{UserRepo: repo}
	return &service
}
