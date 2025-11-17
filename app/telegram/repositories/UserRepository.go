package repositories

import (
	"fmt"
	"strconv"

	"GO/app/core/database"
	"GO/app/telegram/models"
)

type UserRepository struct {
	Repository
	Model models.User
	Db    *database.Mysql
}

func (r *UserRepository) GetModel() models.User {
	return r.Model
}

func (r *UserRepository) GetDB() *database.Mysql {
	return r.Db
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
	model := r.GetModel()
	table := model.GetTable()

	sql := "SELECT * FROM " + table + "WHERE id = " + strconv.Itoa(id)

	row := r.Db.GetDb().QueryRow(sql)
	err := row.Scan(&model.Id)
	if err != nil {
		fmt.Println("ошибка получения юзера")
	}
	return model, err
}
