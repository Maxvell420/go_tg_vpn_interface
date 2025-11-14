package repositories

import (
	"GO/app/core/database"
	"GO/app/telegram/models"
)

type Repository interface {
	GetByID() (models.Model, error)
	GetModel() models.Model
	GetDB() database.Mysql
}
