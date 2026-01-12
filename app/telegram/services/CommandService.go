package services

import (
	"fmt"
	"strings"

	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
	"GO/app/telegram/updates"
)

type CommandService struct {
	UserRepository        *Repositories.UserRepository
	OutworldFacade        *outworld.OutworldFacade
	ReferalLinkRepository *Repositories.ReferalLinkRepository
	ReferalUserRepository *Repositories.ReferalUserRepository
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
		link, err := s.ReferalLinkRepository.GetByHash(hash)
		if err == nil {
			referal_user, err := s.ReferalUserRepository.GetByTgId(link.GetTgId())
			if err != nil || *referal_user.GetID() != update.GetUser() {
				user, err := s.UserRepository.GetByTgID(link.GetTgId())
				if err == nil {
					model := s.ReferalUserRepository.BuildModel(*user.GetTgId(), link.GetTgId())
					_, err = s.ReferalUserRepository.Persist(model)
				}
			}

		}
	}

	user, err := s.UserRepository.GetByTgID(update.GetUser())
	if err != nil {
		fmt.Println(err)
	}
	s.OutworldFacade.SendTelegramStartMessage(*user.GetTgId())
}
