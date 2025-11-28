package services

import (
	"GO/app/libs/telegram"
)

type TelegramBot struct {
	BotToken *string
	Lib      *telegram.Request
}

// Вероятно тут будет логирование
func (s *TelegramBot) SendPost(req telegram.PostRequest) {
	s.Lib.SendPost(req)
}

// Вероятно потом тут будет структура с найстройками
func (s *TelegramBot) BuildMessage(chat_id int, message string) telegram.PostRequest {
	tg_req := telegram.Message{ChatID: chat_id, Text: message}
	return telegram.PostRequest{Method: telegram.SendMessage, Message: &tg_req}
}
