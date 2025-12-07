package updates

import (
	"GO/app/domain/User/Models"
)

type NewChatMember struct {
	Status Models.ChatStatus
	User   User
}
