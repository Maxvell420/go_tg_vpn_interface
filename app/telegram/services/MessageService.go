package services

import (
	"fmt"

	"GO/app/libs/telegram"
	"GO/app/outworld"
	"GO/app/telegram/repositories"
	"GO/app/telegram/updates"
)

type MessageService struct {
	UserRepo       *repositories.UserRepository
	OutworldFacade *outworld.OutworldFacade
}

func (s *MessageService) HandleMessageUpdate(update updates.Message) {
	user_id := update.From.Id
	user, err := s.UserRepo.GetByTgID(user_id)
	if err != nil {
		fmt.Println(err)
	}
	message := telegram.Message{
		ChatID: *user.Tg_id, Text: "hello world",
	}

	data := telegram.PostRequest{
		Message: &message, Method: telegram.SendMessage,
	}

	s.OutworldFacade.SendTelegramPost(data)
}
