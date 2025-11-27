package models

import (
	"database/sql"

	"GO/app/telegram/updates"
)

type User struct {
	Id        *int
	Tg_id     *int
	User_name *string
	Kicked    string
	Is_admin  string
}

func (u *User) FromDB(row *sql.Row) (User, error) {
	var user User
	err := row.Scan(&user.Id, &user.Tg_id, &user.User_name, &user.Kicked, &user.Is_admin)
	return user, err
}

func (u *User) FromData(
	Tg_id int,
	User_name *string,
	Kicked string,
	Is_admin string,
) User {
	user := User{User_name: User_name, Kicked: Kicked, Is_admin: Is_admin, Tg_id: &Tg_id}
	return user
}

func (u *User) UpdateStatus(status updates.ChatStatus) {
	if status == updates.Kicked {
		u.Kicked = "yes"
	} else {
		u.Kicked = "no"
	}
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
	return u.Kicked == "yes"
}

func (u *User) IsAdmin() bool {
	return u.Is_admin == "yes"
}
