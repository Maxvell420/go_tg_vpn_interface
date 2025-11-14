package repositories

import (
	"GO/app/telegram/models"
)

type UserRepository struct {
	Repository
	model models.User
}

func (r *UserRepository) GetModel() models.User {
	var user models.User
	return user
}

func GetByID() (models.User, error) {
}
