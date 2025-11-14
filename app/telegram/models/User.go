package models

type User struct {
	id int
}

func (u *User) GetTable() string {
	return "users"
}

func (u *User) getID() int {
	return u.id
}
