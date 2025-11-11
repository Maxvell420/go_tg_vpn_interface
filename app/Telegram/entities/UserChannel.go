package entities

import (
	"GO/app/Telegram/updates"
)

type UserChannel struct {
	Timestamp int
	Update    updates.TelegramUpdate
}
