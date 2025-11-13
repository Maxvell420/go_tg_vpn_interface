package updates

type MyChatMember struct {
	Chat Chat
	From Chat
}

func (m *MyChatMember) GetUpdateType() UpdateType {
	return NewChatMemberType
}
