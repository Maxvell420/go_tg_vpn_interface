package telegram

type Message struct {
	ChatID              int          `json:"chat_id"`
	Text                string       `json:"text"`
	ParseMode           *string      `json:"parse_mode,omitempty"` // Исправлено имя и добавлен omitempty
	DisableNotification *bool        `json:"disable_notification,omitempty"`
	ProtectContent      *bool        `json:"protect_content,omitempty"`
	Reply_markup        *ReplyMarkup `json:"reply_markup,omitempty"`
}

type KeyboardAction string

const (
	Statistic KeyboardAction = "Посмотреть статистику"
)

type Button struct {
	Text          *string `json:"text,omitempty"`
	Callback_data string  `json:"callback_data,omitempty"`
	Url           *string `json:"url,omitempty"`
}

type ReplyMarkup struct {
	InlineKeyboard          [][]Button `json:"inline_keyboard,omitempty"`
	Remove_keyboard         *bool      `json:"remove_keyboard,omitempty"`
	Force_reply             *bool      `json:"force_reply,omitempty"`
	Selective               *bool      `json:"selective,omitempty"`
	Input_field_placeholder *string    `json:"input_field_placeholder,omitempty"`
}
