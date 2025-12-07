package updates

import "GO/app/domain/User/Models"

type MyChatMember struct {
	Chat            Chat
	From            From
	New_chat_member NewChatMember
}

func (m *MyChatMember) GetUser() int {
	return m.From.Id
}

func (m *MyChatMember) GetNewStatus() Models.ChatStatus {
	return m.New_chat_member.Status
}
