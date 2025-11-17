package models

type User struct {
	Id int
}

func (u User) GetTable() string {
	return "users"
}

func (u *User) getID() int {
	return u.Id
}
