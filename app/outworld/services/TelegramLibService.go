package services

import (
	"GO/app/libs/telegram"
)

type TelegramBot struct {
	BotToken        *string
	Lib             *telegram.Request
	KeyboardService *KeyboardService
}

// Вероятно тут будет логирование
func (s *TelegramBot) sendPost(req telegram.PostRequest) {
	s.Lib.SendPost(req)
}

func (s *TelegramBot) SendTelegramStartMessage(chat_id int) {
	keyboard := s.KeyboardService.GetStartKeyboard()
	message := telegram.Message{ChatID: chat_id, Text: "Приветствую", Reply_markup: &keyboard}
	req := telegram.PostRequest{Method: telegram.SendMessage, Message: &message}
	s.sendPost(req)
}
