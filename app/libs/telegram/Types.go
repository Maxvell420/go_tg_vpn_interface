package telegram

type Message struct {
	ChatID              int     `json:"chat_id"`
	Text                string  `json:"text"`
	ParseMode           *string `json:"parse_mode,omitempty"` // Исправлено имя и добавлен omitempty
	DisableNotification *bool   `json:"disable_notification,omitempty"`
	ProtectContent      *bool   `json:"protect_content,omitempty"`
}
