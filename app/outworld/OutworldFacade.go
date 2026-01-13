package outworld

import (
	"GO/app/core"
	"GO/app/libs/3xui"
	"GO/app/libs/telegram"
	"GO/app/outworld/services"
)

// Фасад который ходит во внешний мир,через него собирается структура для отправки в телегу
type OutworldFacade struct {
	Cntx *core.Context
}

// уйти от этого метода
func (f *OutworldFacade) SendTelegramPost(req telegram.PostRequest) {
	service := f.buildTelegramBotService()
	service.Lib.SendPost(req)
}

func (f *OutworldFacade) SendTelegramStartMessage(chat_id int) {
	service := f.buildTelegramBotService()
	service.SendTelegramStartMessage(chat_id)
}

func (f *OutworldFacade) SendTelegramRefLinkMessage(chat_id int, link string) {
	service := f.buildTelegramBotService()
	service.SendTelegramRefLinkMessage(chat_id, link)
}

func (f *OutworldFacade) GetInbounds(chat_id int) {
	manager := f.buildOutworldManager()
	manager.GetInbounds(chat_id)
}

func (f *OutworldFacade) buildOutworldManager() services.OutworldManager {
	return services.OutworldManager{TelegramBot: f.buildTelegramBotService(), XuiService: f.buildXuiService()}
}

func (f *OutworldFacade) buildXuiService() *services.XuiService {
	return &services.XuiService{Request: f.buildXuiLib()}
}

func (f *OutworldFacade) buildTelegramBotService() *services.TelegramBot {
	return &services.TelegramBot{Lib: f.buildTelegramLib(), KeyboardService: f.buildKeyboardService()}
}

func (f *OutworldFacade) buildTelegramLib() *telegram.Request {
	return f.Cntx.GetTelegramRequest()
}

func (f *OutworldFacade) buildKeyboardService() *services.KeyboardService {
	return &services.KeyboardService{}
}

func (f *OutworldFacade) buildXuiLib() *xui.Request {
	return f.Cntx.Get3xuiRequest()
}
