package database

import (
	"database/sql"
)

type Model interface {
	GetID() *int
	GetTable() string
	FromDB(row *sql.Row) (Model, error)
}
