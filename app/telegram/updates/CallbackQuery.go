package updates

type CallbackQuery struct {
	Id      string
	From    From
	Message Message
	Chat    Chat
	Data    string
}

func (m *CallbackQuery) GetUser() int {
	return m.Message.GetUser()
}

type CallbackAction string

const (
	Inbounds CallbackAction = "Inbounds"
)
