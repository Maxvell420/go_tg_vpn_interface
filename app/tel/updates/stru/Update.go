package stru

import (
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
