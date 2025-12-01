package entities

import (
	"GO/app/telegram/updates"
)

type UserChannel struct {
	Timestamp int
	Update    *updates.Update
	Ch        *chan *updates.Update
}

