package services

import (
	"fmt"
	"strings"

	"GO/app/domain/User/Repositories"
	"GO/app/outworld"
	"GO/app/telegram/entities"
	"GO/app/telegram/updates"
)

type CommandService struct {
	UserRepository *Repositories.UserRepository
	OutworldFacade *outworld.OutworldFacade
	ReferalService *ReferalService
}

func (s *CommandService) HandleCommand(update updates.Message, jobsChannel chan entities.Job) {
	// тут будет парсинг команд
	if strings.HasPrefix(*update.Text, string(updates.Start)) {
		s.HandleStartCommand(update)
	}

	if *update.Text == string(updates.RefLink) {
		s.HandleRefLinkCommand(update, jobsChannel)
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

func (s *CommandService) HandleRefLinkCommand(update updates.Message, jobsChannel chan entities.Job) {
	link := s.ReferalService.GetUserRefLink(update.GetUser())
	jobsChannel <- entities.Job{Type: entities.TrafficUsage, Data: &map[string]string{"link": link}, UserId: update.GetUser()}
	s.OutworldFacade.SendTelegramRefLinkMessage(update.GetUser(), link)
}
