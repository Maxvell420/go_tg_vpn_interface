package stru

type UserChannel struct {
	Update    Update
	Timestamp int
	Channel   *chan any
}
