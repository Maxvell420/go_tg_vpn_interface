package updates

type TelegramUpdate struct {
	Ok     bool
	Result Update
}

type UserUpdate interface {
	GetUpdateId
	GetUpdateType
	GetMyChatMember
	GetMessage
	GetUserId
}

type UpdateType int

const (
	MessageType      UpdateType = 1
	MyChatMemberType UpdateType = 2
)

type GetUpdateType interface {
	GetUpdateType() UpdateType
}

type GetUpdateId interface {
	GetUpdateId() int
}

type GetMessage interface {
	GetMessage() *Message
}

type GetMyChatMember interface {
	GetMyChatMember() *MyChatMember
}

type GetUserId interface {
	GetUserId() int
}
