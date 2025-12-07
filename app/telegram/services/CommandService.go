package services

import (
	"fmt"

	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
	"GO/app/telegram/updates"
)

type CommandService struct {
	UserRepository *Repositories.UserRepository
	OutworldFacade *outworld.OutworldFacade
}

func (s *CommandService) HandleCommand(update updates.Message) {
	// тут будет парсинг команд
	user, err := s.UserRepository.GetByTgID(update.GetUser())
	if err != nil {
		fmt.Println(err)
	}
	s.OutworldFacade.SendTelegramStartMessage(*user.GetTgId())
}
