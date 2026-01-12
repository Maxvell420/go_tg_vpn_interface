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
	for _, entity := range m.Entities {
		if entity.Type == "bot_command" {
			return true
		}
	}
	return false
}

type ChatCommand string

const (
	Start ChatCommand = "/start"
)
