package stru

import (
	"GO/app/tel/updates/interfaces"
	"GO/app/tel/updates/stru/up_types"
)

type Update struct {
	Ok     bool
	Result []Result
}

type Result struct {
	Update_id    int
	Message      *up_types.Message
	MyChatMember *up_types.MyChatMember
}

func (r *Result) GetUpdateId() int {
	return r.Update_id
}

func (r *Result) GetUpdateType() interfaces.UpdateType {
	if r.Message != nil {
		return interfaces.Message
	}

	if r.MyChatMember != nil {
		return interfaces.NewChatMember
	}
	panic(1)
	// return error
}
