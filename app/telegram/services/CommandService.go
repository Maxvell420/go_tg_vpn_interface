package services

import (
	"fmt"
	"strings"

	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
	"GO/app/telegram/updates"
)

type CommandService struct {
	UserRepository *Repositories.UserRepository
	OutworldFacade *outworld.OutworldFacade
	ReferalService *ReferalService
}

func (s *CommandService) HandleCommand(update updates.Message) {
	// тут будет парсинг команд
	if strings.HasPrefix(*update.Text, string(updates.Start)) {
		s.HandleStartCommand(update)
	}
}

func (s *CommandService) HandleStartCommand(update updates.Message) {
	if *update.Text != string(updates.Start) {
		hash := strings.Replace(*update.Text, string(updates.Start)+" ", "", 1)
		s.ReferalService.HandleStartReferal(update.GetUser(), hash)
	}

	user, err := s.UserRepository.GetByTgID(update.GetUser())
	if err != nil {
		fmt.Println(err)
	}
	s.OutworldFacade.SendTelegramStartMessage(*user.GetTgId())
}
