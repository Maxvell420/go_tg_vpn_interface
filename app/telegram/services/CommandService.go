package services

import (
	"strings"

	"GO/app/domain/User"
	"GO/app/outworld"
	"GO/app/telegram/entities"
	"GO/app/telegram/updates"
)

type CommandService struct {
	OutworldFacade *outworld.OutworldFacade
	UserFacade     *user.UserFacade
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
		s.UserFacade.HandleStartReferal(update.GetUser(), hash)
	}

	s.OutworldFacade.SendTelegramStartMessage(update.GetUser())
}

func (s *CommandService) HandleRefLinkCommand(update updates.Message, jobsChannel chan entities.Job) {
	link := s.UserFacade.GetUserRefLink(update.GetUser())
	s.OutworldFacade.SendTelegramRefLinkMessage(update.GetUser(), link)
}
