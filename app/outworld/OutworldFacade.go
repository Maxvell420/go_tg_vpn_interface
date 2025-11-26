package outworld

import (
	"GO/app/core"
	"GO/app/libs/telegram"
	"GO/app/outworld/services"
)

type OutworldFacade struct {
	Cntx *core.Context
}

// Фасад который ходит во внешний мир
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
