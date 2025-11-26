package repositories

import (
	"database/sql"
	"fmt"
	"strconv"

	"GO/app/telegram/models"
)

type UserRepository struct {
	Repository
	Model models.User
	Db    *sql.DB
}

func (r *UserRepository) GetModel() models.User {
	return r.Model
}

func (r *UserRepository) GetDB() *sql.DB {
	return r.Db
}

func (r *UserRepository) GetByID(id int) *models.User {
	model := r.GetModel()
	table := model.GetTable()
	sql := "SELECT id,tg_id,user_name,kicked,is_admin FROM " + table + " WHERE id = " + strconv.Itoa(id)
	row := r.Db.QueryRow(sql)
	user, err := model.FromDB(row)
	if err != nil {
		fmt.Println("ошибка получения юзера")
	}
	return &user
}

func (r *UserRepository) GetByTgID(id int) (*models.User, error) {
	model := r.GetModel()
	table := model.GetTable()

	sql := "SELECT id,tg_id,user_name,kicked,is_admin FROM " + table + " WHERE tg_id = " + strconv.Itoa(id)
	row := r.Db.QueryRow(sql)
	user, err := model.FromDB(row)
	if err != nil {
		fmt.Println("ошибка получения юзера")
	}
	return &user, err
}
