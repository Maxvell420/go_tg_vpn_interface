package database

type Repository interface {
	GetByID() (Model, error)
	GetModel() Model
	GetDB() Mysql
	BuildModel() Model
}
