package outworld

import (
	"GO/app/core"
	"GO/app/libs/telegram"
	"GO/app/outworld/services"
)

// Фасад который ходит во внешний мир,через него собирается структура для отправки в телегу
type OutworldFacade struct {
	Cntx *core.Context
}

func (f *OutworldFacade) SendTelegramPost(req telegram.PostRequest) {
	service := f.buildTelegramBotService()
	service.Lib.SendPost(req)
}

func (f *OutworldFacade) SendTelegramStartMessage(chat_id int) {
	service := f.buildTelegramBotService()
	service.SendTelegramStartMessage(chat_id)
}

func (f *OutworldFacade) buildTelegramBotService() services.TelegramBot {
	token := f.Cntx.GetSecrets().BotToken
	return services.TelegramBot{BotToken: token, Lib: f.buildTelegramLib(), KeyboardService: f.buildKeyboardService()}
}

func (f *OutworldFacade) buildTelegramLib() *telegram.Request {
	token := f.Cntx.GetSecrets().BotToken
	return &telegram.Request{BotToken: token}
}

func (f *OutworldFacade) buildKeyboardService() *services.KeyboardService {
	return &services.KeyboardService{}
}
