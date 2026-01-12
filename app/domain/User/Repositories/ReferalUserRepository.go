package Repositories

import (
	"database/sql"

	"GO/app/core/database"
	"GO/app/domain/User/Models"
)

type ReferalUserRepository struct {
	database.Repository
	Model *Models.ReferalUser
	Db    *sql.DB
}

func (r *ReferalUserRepository) GetModel() *Models.ReferalUser {
	return r.Model
}

func (r *ReferalUserRepository) Persist(user Models.ReferalUser) (Models.ReferalUser, error) {
	var sql string
	var err error

	if user.GetID() != nil {
		sql = "UPDATE referal_users SET tg_id = ?, owner_tg_id = ? WHERE id = ?"
		_, err = r.Db.Exec(sql, user.GetTgId(), user.GetOwnerTgId(), user.GetID())
	} else {
		sql = "INSERT INTO referal_users(tg_id, owner_tg_id) VALUES (?, ?)"
		_, err = r.Db.Exec(sql, user.GetTgId(), user.GetOwnerTgId())
	}

	return user, err
}

func (r *ReferalUserRepository) GetByTgId(tg_id int) (Models.ReferalUser, error) {
	var sql string
	var err error

	sql = "SELECT id, tg_id, owner_tg_id FROM referal_users WHERE tg_id = ?"
	row := r.Db.QueryRow(sql, tg_id)
	user, err := r.buildReferalUserModel(row)
	return user, err
}

func (r *ReferalUserRepository) buildReferalUserModel(row *sql.Row) (Models.ReferalUser, error) {
	var user Models.ReferalUser
	err := row.Scan(&user.Id, &user.Tg_id, &user.Owner_tg_id)
	return user, err
}

func (r *ReferalUserRepository) BuildModel(tg_id int, owner_tg_id int) Models.ReferalUser {
	return Models.ReferalUser{Tg_id: &tg_id, Owner_tg_id: &owner_tg_id}
}
