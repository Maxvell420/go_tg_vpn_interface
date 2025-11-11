package updates

type TelegramUpdate interface {
	GetUpdateId
	GetUpdateType
	GetMessage
	GetNewChatMember
}

type UpdateType int

const (
	MessageType       UpdateType = 1
	NewChatMemberType UpdateType = 2
)

type GetUpdateType interface {
	GetUpdateType() UpdateType
}

type GetUpdateId interface {
	GetUpdateId() int
}

type GetMessage interface {
	getMessage() *Message
}

type GetNewChatMember interface {
	GetNewChatMember() *NewChatMember
}
