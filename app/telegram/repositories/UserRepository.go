package repositories

import (
	"database/sql"
	"fmt"
	"strconv"

	"GO/app/telegram/models"

	"github.com/davecgh/go-spew/spew"
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

func (r *UserRepository) GetByID(id int) models.User {
	model := r.GetModel()
	table := model.GetTable()
	sql := "SELECT id,tg_id,user_name,kicked,is_admin FROM " + table + " WHERE id = " + strconv.Itoa(id)
	row := r.Db.QueryRow(sql)
	user, err := model.FromDB(row)
	if err != nil {
		fmt.Println("ошибка получения юзера")
	}
	return user
}

func (r *UserRepository) GetByTgID(id int) (models.User, error) {
	model := r.GetModel()
	table := model.GetTable()

	sql := "SELECT id,tg_id,user_name,kicked,is_admin FROM " + table + " WHERE tg_id = " + strconv.Itoa(id)
	row := r.Db.QueryRow(sql)
	user, err := model.FromDB(row)
	return user, err
}

func (r *UserRepository) Persist(user models.User) {
	var sql string
	var err error

	if user.GetID() != nil {
		sql = "UPDATE users SET tg_id = ?, user_name = ?, kicked = ?, is_admin = ? WHERE id = ?"
		_, err = r.Db.Exec(sql, user.Tg_id, user.User_name, user.Kicked, user.Is_admin, user.Id)
	} else {
		sql = "INSERT INTO users(tg_id, user_name, kicked, is_admin) VALUES (?, ?, ?, ?)"
		_, err = r.Db.Exec(sql, user.GetTgId(), user.User_name, user.Kicked, user.Is_admin)
	}
	if err != nil {
		spew.Dump(err)
	}
}
