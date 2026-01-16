package Repositories

import (
	"database/sql"

	"GO/app/core/database"
	"GO/app/domain/User/Models"
)

type InboundRepository struct {
	database.Repository
	Model *Models.Inbound
	Db    *sql.DB
}

func (r *InboundRepository) GetModel() *Models.Inbound {
	return r.Model
}

func (r *InboundRepository) GetDB() *sql.DB {
	return r.Db
}
