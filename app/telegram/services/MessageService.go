package services

import (
	"fmt"

	"GO/app/domain/User/Repositories"
	"GO/app/libs/telegram"
	"GO/app/outworld"
	"GO/app/telegram/updates"
)

type MessageService struct {
	UserRepo        *Repositories.UserRepository
	OutworldFacade  *outworld.OutworldFacade
	CommandsHandler *CommandService
}

func (s *MessageService) HandleMessageUpdate(update updates.Message) {
	// Думаю тут будет поиск стейтов
	if update.IsCommand() {
		s.CommandsHandler.HandleCommand(update)
	} else {
		s.HandleRegularMessage(update)
	}
}

func (s *MessageService) HandleRegularMessage(update updates.Message) {
	user_id := update.From.Id
	user, err := s.UserRepo.GetByTgID(user_id)
	if err != nil {
		fmt.Println(err)
	}

	// Перенести это все в outworld
	message := telegram.Message{
		ChatID: *user.GetTgId(), Text: "Это текст заглушка",
	}

	data := telegram.PostRequest{
		Message: &message, Method: telegram.SendMessage,
	}

	s.OutworldFacade.SendTelegramPost(data)
}
