package updates

type NewChatMember struct {
	Status ChatStatus
	User   User
}

type ChatStatus string

const (
	Administrator ChatStatus = "administrator"
	Member        ChatStatus = "member"
	Restricted    ChatStatus = "restricted"
	Left          ChatStatus = "left"
	Kicked        ChatStatus = "kicked"
)
