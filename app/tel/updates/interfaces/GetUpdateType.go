package interfaces

type GetUpdateType interface {
	GetUpdateType() UpdateType
}

type UpdateType int

const (
	Message       UpdateType = 1
	NewChatMember UpdateType = 2
)
