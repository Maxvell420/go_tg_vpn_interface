package services

import "GO/app/libs/3xui"

// Этот сервис отвечает за управление внешним миром когда нужно связать несколько сервисов вместе
type OutworldManager struct {
	TelegramBot *TelegramBot
	XuiService  *XuiService
}

func (m *OutworldManager) GetInbounds() []xui.ListObj {
	return m.XuiService.GetInbounds()
}
