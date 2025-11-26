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
