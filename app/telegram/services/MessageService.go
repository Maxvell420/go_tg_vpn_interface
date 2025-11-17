package services

import (
	"fmt"

	"GO/app/telegram/repositories"
	"GO/app/telegram/updates"
)

type MessageService struct {
	UserRepo *repositories.UserRepository
}

func (s *MessageService) HandleMessageUpdate(update updates.Message) {
	user_id := update.From.Id
	user, err := s.UserRepo.GetByTgID(user_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}
