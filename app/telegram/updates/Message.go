package updates

type Message struct {
	Message_id int
	From       From
	Chat       Chat
	Entities   []Entity
	Date       int
	Text       *string
}

func (m *Message) GetUser() int {
	return m.From.Id
}

func (m *Message) IsCommand() bool {
	switch *m.Text {
	case string(Start):
		return true
	}
	return false
}

type ChatCommand string

const (
	Start ChatCommand = "/start"
)
