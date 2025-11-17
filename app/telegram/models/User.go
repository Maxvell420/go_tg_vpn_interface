package models

type User struct {
	Id    int
	Tg_id int
}

func (u User) GetTable() string {
	return "users"
}

func (u *User) GetID() int {
	return u.Id
}

func (u *User) GetTgId() int {
	return u.Tg_id
}
