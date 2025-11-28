package outworld

import (
	"GO/app/core"
	"GO/app/libs/telegram"
	"GO/app/outworld/services"
)

// Фасад который ходит во внешний мир
type OutworldFacade struct {
	Cntx *core.Context
}

func (f *OutworldFacade) SendTelegramPost(req telegram.PostRequest) {
	service := f.buildTelegramBotService()
	service.Lib.SendPost(req)
}

func (f *OutworldFacade) buildTelegramBotService() services.TelegramBot {
	token := f.Cntx.GetSecrets().BotToken
	return services.TelegramBot{BotToken: token, Lib: f.buildTelegramLib()}
}

func (f *OutworldFacade) buildTelegramLib() *telegram.Request {
	token := f.Cntx.GetSecrets().BotToken
	return &telegram.Request{BotToken: token}
}

// Единая точка для формирования сообщений в либу телеги
func (f *OutworldFacade) BuildTelegramMessage(chat_id int, message string) telegram.PostRequest {
	service := f.buildTelegramBotService()
	return service.BuildMessage(chat_id, message)
}
