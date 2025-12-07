package Models

type User struct {
	Id        *int
	Tg_id     *int
	User_name *string
	Kicked    string
	Is_admin  string
}

type ChatStatus string

const (
	Administrator ChatStatus = "administrator"
	Member        ChatStatus = "member"
	Restricted    ChatStatus = "restricted"
	Left          ChatStatus = "left"
	Kicked        ChatStatus = "kicked"
)

func (u *User) FromData(
	Tg_id int,
	User_name *string,
	Kicked string,
	Is_admin string,
) User {
	user := User{User_name: User_name, Kicked: Kicked, Is_admin: Is_admin, Tg_id: &Tg_id}
	return user
}

func (u *User) UpdateStatus(status ChatStatus) {
	if status == Kicked {
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

func (u *User) GetUserName() *string {
	return u.User_name
}

func (u *User) GetKicked() string {
	return u.Kicked
}

func (u *User) GetAdmin() string {
	return u.Is_admin
}
