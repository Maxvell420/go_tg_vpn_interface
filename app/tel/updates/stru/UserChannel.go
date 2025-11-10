package stru

type UserChannel struct {
	update    Update
	timestamp int
	channel   *chan any
}
