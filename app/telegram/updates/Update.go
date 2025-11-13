package updates

type TelegramUpdate struct {
	Ok     bool
	Result Update
}

type UserUpdate interface {
	GetUpdateId
	GetUpdateType
	// GetMessage
	// GetNewChatMember
}

type Update struct {
	Update_id int
	Message   *Message
	// Добавить сюда все остальные типы
}

func (u *Update) GetUpdateType() UpdateType {
	if u.Message != nil {
		return MessageType
	}

	// Добавить все остальные
	return NewChatMemberType
}

func (u *Update) GetUpdateId() int {
	return u.Update_id
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
