package services

import (
	"fmt"

	"GO/app/domain/User"
	"GO/app/libs/telegram"
	"GO/app/outworld"
	"GO/app/telegram/entities"
	"GO/app/telegram/updates"
)

// TODO: Вообще если логика будет сложнее то лучше будет вынести это в отдельный сервис как и все сервисы телеграмма и избавиться от параметров функций в виде апдейтов телеги т.к. это нарушение направления домена
// TODO: Добавить логику для обработки обычных сообщений
type MessageService struct {
	UserFacade      *user.UserFacade
	OutworldFacade  *outworld.OutworldFacade
	CommandsHandler *CommandService
}

func (s *MessageService) HandleMessageUpdate(update updates.Message, jobsChannel chan entities.Job) {
	// Думаю тут будет поиск стейтов
	if update.IsCommand() {
		s.CommandsHandler.HandleCommand(update, jobsChannel)
	} else {
		s.HandleRegularMessage(update)
	}
}

func (s *MessageService) HandleRegularMessage(update updates.Message) {
	user_id := update.From.Id
	user, err := s.UserFacade.GetUserByTgId(user_id)
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
