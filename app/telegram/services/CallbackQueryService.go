package services

import (
	"strings"

	"GO/app/outworld"
	"GO/app/telegram/updates"
)

type CallbackQueryService struct {
	OutworldFacade *outworld.OutworldFacade
}

func (s *CallbackQueryService) HandleCallbackQuery(update updates.CallbackQuery) {
	parts := strings.Split(update.Data, ";")

	var action, data string

	for _, part := range parts {
		// Разделяем каждую часть по двоеточию
		keyValue := strings.SplitN(part, ":", 2)
		if len(keyValue) == 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])

			switch key {
			case "action":
				action = value
			case "data":
				data = value
			}
		}
	}

	switch action {
	case string(updates.Inbounds):
		s.HandleInbounds(update.GetUser(), data)
	default:
		panic(1)
	}
}

func (s *CallbackQueryService) HandleInbounds(chat_id int, data string) {
	s.OutworldFacade.GetInbounds(chat_id)
}
