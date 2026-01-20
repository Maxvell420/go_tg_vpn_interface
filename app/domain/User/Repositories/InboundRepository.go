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

func (r *InboundRepository) GetInbounds() map[int]Models.Inbound {
	model := r.GetModel()
	table := model.GetTable()
	sql := "SELECT id,total,calc_total,protocol,tag FROM " + table
	rows, err := r.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	inbounds := make(map[int]Models.Inbound, 0)
	for rows.Next() {
		var inbound Models.Inbound
		err = rows.Scan(&inbound.Id, &inbound.Total, &inbound.CalcTotal, &inbound.Protocol, &inbound.Tag)
		if err != nil {
			panic(err)
		}
		inbounds[*inbound.Id] = inbound
	}
	return inbounds
}
