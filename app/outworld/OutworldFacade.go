package outworld

import (
	"GO/app/libs/3xui"
	"GO/app/libs/telegram"
)

// Фасад который ходит во внешний мир,через него собирается структура для отправки в телегу
type OutworldFacade struct {
	Builder *Builder
}

// уйти от этого метода
func (f *OutworldFacade) SendTelegramPost(req telegram.PostRequest) {
	service := f.Builder.BuildTelegramBotService()
	service.Lib.SendPost(req)
}

func (f *OutworldFacade) SendTelegramStartMessage(chat_id int) {
	service := f.Builder.BuildTelegramBotService()
	service.SendTelegramStartMessage(chat_id)
}

func (f *OutworldFacade) SendTelegramRefLinkMessage(chat_id int, link string) {
	service := f.Builder.BuildTelegramBotService()
	service.SendTelegramRefLinkMessage(chat_id, link)
}

func (f *OutworldFacade) GetInbounds() []xui.ListObj {
	manager := f.Builder.BuildOutworldManager()
	return manager.GetInbounds()
}
