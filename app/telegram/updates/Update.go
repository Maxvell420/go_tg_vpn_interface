package updates

type Update struct {
	Update_id      int
	Message        *Message
	My_chat_member *MyChatMember
	Callback_query *CallbackQuery
	// Добавить сюда все остальные типы
}

func (u *Update) GetUpdateType() UpdateType {
	if u.Message != nil {
		return MessageType
	} else if u.My_chat_member != nil {
		return MyChatMemberType
	} else if u.Callback_query != nil {
		return CallbackQueryType
	}

	// Добавить все остальные
	return MyChatMemberType
}

func (u *Update) GetUpdateId() int {
	return u.Update_id
}

func (u *Update) GetMessage() *Message {
	return u.Message
}

func (u *Update) GetMyChatMember() *MyChatMember {
	return u.My_chat_member
}

func (u *Update) GetCallbackQuery() *CallbackQuery {
	return u.Callback_query
}

func (u *Update) GetUserId() int {
	update_type := u.GetUpdateType()

	switch update_type {
	case MessageType:
		return u.Message.GetUser()
	case MyChatMemberType:
		return u.My_chat_member.GetUser()
	case CallbackQueryType:
		return u.Callback_query.GetUser()
	}
	panic(1)
}
