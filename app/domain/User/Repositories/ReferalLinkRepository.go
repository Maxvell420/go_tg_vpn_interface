package Repositories

import (
	"database/sql"

	"GO/app/core/database"
	"GO/app/domain/User/Models"
)

type ReferalLinkRepository struct {
	database.Repository
	Model *Models.ReferalLink
	Db    *sql.DB
}

func (r *ReferalLinkRepository) GetModel() *Models.ReferalLink {
	return r.Model
}

func (r *ReferalLinkRepository) Persist(link Models.ReferalLink) (Models.ReferalLink, error) {
	var sql string
	var err error

	if r.Model.GetID() != nil {
		sql = "UPDATE referal_links SET hash = ?, tg_id = ? WHERE id = ?"
		_, err = r.Db.Exec(sql, link.GetHash(), link.GetTgId(), r.Model.GetID())
	} else {
		sql = "INSERT INTO referal_links(hash, tg_id) VALUES (?, ?)"
		_, err = r.Db.Exec(sql, link.GetHash(), link.GetTgId())
	}

	return link, err
}

func (r *ReferalLinkRepository) GetByHash(hash string) (Models.ReferalLink, error) {
	var sql string
	var err error

	sql = "SELECT id, hash, tg_id FROM referal_links WHERE hash = ?"
	row := r.Db.QueryRow(sql, hash)
	link, err := r.buildReferalLinkModel(row)

	return link, err
}

func (r *ReferalLinkRepository) GetByTgId(tg_id int) (Models.ReferalLink, error) {
	var sql string
	var err error

	sql = "SELECT id, hash, tg_id FROM referal_links WHERE tg_id = ?"
	row := r.Db.QueryRow(sql, tg_id)
	link, err := r.buildReferalLinkModel(row)
	return link, err
}

func (r *ReferalLinkRepository) BuildModel(hash string, tg_id int) Models.ReferalLink {
	return Models.ReferalLink{Hash: &hash, Tg_id: &tg_id}
}

func (r *ReferalLinkRepository) buildReferalLinkModel(row *sql.Row) (Models.ReferalLink, error) {
	var link Models.ReferalLink
	err := row.Scan(&link.Id, &link.Hash, &link.Tg_id)
	return link, err
}
