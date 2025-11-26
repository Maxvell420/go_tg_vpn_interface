package models

import (
	"database/sql"
)

type User struct {
	Id        *int
	Tg_id     *int
	User_name *string
	Kicked    *string
	Is_admin  *string
}

func (u *User) FromDB(row *sql.Row) (User, error) {
	var user User
	err := row.Scan(&user.Id, &user.Tg_id, &user.User_name, &user.Kicked, &user.Is_admin)
	return user, err
}

func (u *User) GetTable() string {
	return "users"
}

func (u *User) GetID() *int {
	return u.Id
}

func (u *User) GetTgId() *int {
	return u.Tg_id
}

func (u *User) IsKicked() bool {
	return *u.Kicked == "yes"
}

func (u *User) IsAdmin() bool {
	return *u.Is_admin == "yes"
}
