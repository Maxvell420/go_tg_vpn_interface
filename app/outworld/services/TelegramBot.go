package services

import (
	"os"

	"GO/app/libs/telegram"
)

// Этот сервис отвечает за отправку сообщений в телегу
type TelegramBot struct {
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

func (s *TelegramBot) SendTelegramRefLinkMessage(chat_id int, link string) {
	message := telegram.Message{ChatID: chat_id, Text: "Ваша реферальная ссылка: \n" + "https://t.me/" + os.Getenv("bot_name") + "?start=" + link}
	req := telegram.PostRequest{Method: telegram.SendMessage, Message: &message}
	s.sendPost(req)
}

// Реализовать функцию прячущую в телегу сообщение с inline keyboard
