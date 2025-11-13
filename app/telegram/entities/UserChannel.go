package entities

import (
	"GO/app/telegram/updates"
)

type UserChannel struct {
	Timestamp int
	Update    updates.UserUpdate
	Ch        *chan updates.UserUpdate
}
