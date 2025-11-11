package updates

type UserChannel struct {
	Update    Update
	Timestamp int
	Channel   *chan any
}
