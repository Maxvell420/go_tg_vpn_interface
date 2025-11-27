package updates

type MyChatMember struct {
	Chat            Chat
	From            From
	New_chat_member NewChatMember
}

func (m *MyChatMember) GetUser() int {
	return m.From.Id
}

func (m *MyChatMember) GetNewStatus() ChatStatus {
	return m.New_chat_member.Status
}
