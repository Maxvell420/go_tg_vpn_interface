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

func (r *UserRepository) GetByID(id int) (models.User, error) {
	model := r.GetModel()
	table := model.GetTable()

	sql := "SELECT id,tg_id FROM " + table + " WHERE id = " + strconv.Itoa(id)
	row := r.Db.QueryRow(sql)
	err := row.Scan(&model.Id, &model.Tg_id)
	if err != nil {
		fmt.Println("ошибка получения юзера")
	}
	return model, err
}

func (r *UserRepository) GetByTgID(id int) (models.User, error) {
	model := r.GetModel()
	table := model.GetTable()

	sql := "SELECT id,tg_id FROM " + table + " WHERE tg_id = " + strconv.Itoa(id)
	fmt.Println(sql)
	row := r.Db.QueryRow(sql)

	err := row.Scan(&model.Id, &model.Tg_id)
	if err != nil {
		fmt.Println("ошибка получения юзера")
	}
	return model, err
}
