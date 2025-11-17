package services

import (
	"GO/app/telegram/repositories"
	"GO/app/telegram/updates"
)

type MessageService struct {
	UserRepo *repositories.UserRepository
}

func (s *MessageService) HandleMessageUpdate(update updates.Message) {
	user_id := update.From.Id

	user := 1
}
