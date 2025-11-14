package repositories

import (
	"GO/app/telegram/models"
)

type Repository interface {
	GetByID() (models.Model, error)
	GetModel() models.Model
}
