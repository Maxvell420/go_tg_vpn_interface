package repositories

import (
	"GO/app/core/database"
	"GO/app/telegram/models"
)

type UserRepository struct {
	Repository
	model models.User
	db    *database.Mysql
}

func (r *UserRepository) GetModel() *models.User {
	return &r.model
}

func (r *UserRepository) GetDB() *database.Mysql {
	return r.db
}

func (r *UserRepository) GetByID() (models.User, error) {
	table := r.GetModel().GetTable()

	sql := "SELECT * FROM " + table
}
