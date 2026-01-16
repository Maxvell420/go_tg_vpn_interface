package Repositories

import (
	"database/sql"
	"strconv"

	"GO/app/core/database"
	"GO/app/domain/User/Models"
)

type UserRepository struct {
	database.Repository
	// Я достаю по pointer из-за того что мне нужно достать только таблицу из модели
	Model *Models.User
	Db    *sql.DB
}

func (r *UserRepository) GetModel() *Models.User {
	return r.Model
}

func (r *UserRepository) GetDB() *sql.DB {
	return r.Db
}

func (r *UserRepository) GetByID(id int) (Models.UserModel, error) {
	model := r.GetModel()
	table := model.GetTable()
	sql := "SELECT id,tg_id,user_name,kicked,is_admin FROM " + table + " WHERE id = " + strconv.Itoa(id)
	row := r.Db.QueryRow(sql)
	user, err := r.buildUserModel(row)

	return user, err
}

func (r *UserRepository) GetByTgID(id int) (Models.UserModel, error) {
	model := r.GetModel()
	table := model.GetTable()

	sql := "SELECT id,tg_id,user_name,kicked,is_admin FROM " + table + " WHERE tg_id = " + strconv.Itoa(id)
	row := r.Db.QueryRow(sql)
	user, err := r.buildUserModel(row)

	return user, err
}

func (r *UserRepository) buildUserModel(row *sql.Row) (Models.UserModel, error) {
	var user Models.User
	err := row.Scan(&user.Id, &user.Tg_id, &user.User_name, &user.Kicked, &user.Is_admin)
	return &user, err
}

func (r *UserRepository) Persist(user Models.UserModel) {
	var sql string
	var err error

	if user.GetID() != nil {

		sql = "UPDATE users SET tg_id = ?, user_name = ?, kicked = ?, is_admin = ? WHERE id = ?"
		_, err = r.Db.Exec(sql, user.GetTgId(), user.GetUserName(), user.GetKicked(), user.GetAdmin(), user.GetID())
	} else {
		sql = "INSERT INTO users(tg_id, user_name, kicked, is_admin) VALUES (?, ?, ?, ?)"
		_, err = r.Db.Exec(sql, user.GetTgId(), user.GetUserName(), user.GetKicked(), user.GetAdmin())
	}
	if err != nil {
		// TODO: обработать ошибку
	}
}
