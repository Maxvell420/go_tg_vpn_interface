package updates

import (
	"GO/app/tel/updates/struct/up_types"
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
