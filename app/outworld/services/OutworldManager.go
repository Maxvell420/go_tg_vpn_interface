package services

// Этот сервис отвечает за управление внешним миром когда нужно связать несколько сервисов вместе
type OutworldManager struct {
	TelegramBot *TelegramBot
	XuiService  *XuiService
}

func (m *OutworldManager) GetInbounds(chat_id int) {
	m.XuiService.GetInbounds()
}
