package updates

type CallbackQuery struct {
	Id            int
	From          From
	Message       Message
	Chat_instance Chat
	Data          string
}

func (m *CallbackQuery) GetUser() int {
	return m.Message.GetUser()
}

type CallbackAction string

const (
	Inbounds CallbackAction = "Inbounds"
)
