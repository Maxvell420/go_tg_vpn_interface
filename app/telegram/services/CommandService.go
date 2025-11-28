package services

import (
	"fmt"

	"GO/app/outworld"
	"GO/app/telegram/repositories"
	"GO/app/telegram/updates"

	"github.com/davecgh/go-spew/spew"
)

type CommandService struct {
	UserRepository *repositories.UserRepository
	OutworldFacade *outworld.OutworldFacade
}

func (s *CommandService) HandleCommand(update updates.Message) {
	spew.Dump(update.GetUser())
	user, err := s.UserRepository.GetByTgID(update.GetUser())
	if err != nil {
		fmt.Println(err)
	}

	post := s.OutworldFacade.BuildTelegramMessage(*user.Tg_id, "Стартовое сообщение")
	s.OutworldFacade.SendTelegramPost(post)
}
