package Models

type UserModel interface {
	GetTable() string
	GetID() *int
	GetTgId() *int
	IsKicked() bool
	IsAdmin() bool
	GetUserName() *string
	GetKicked() string
	GetAdmin() string
	UpdateStatus(status ChatStatus)
}
