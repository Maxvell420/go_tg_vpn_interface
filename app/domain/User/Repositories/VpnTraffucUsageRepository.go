package Repositories

import (
	"database/sql"

	"GO/app/core/database"
	"GO/app/domain/User/Models"
)

type VpnTrafficUsageRepository struct {
	database.Repository
	Model *Models.VpnTrafficUsage
	Db    *sql.DB
}

func (r *VpnTrafficUsageRepository) GetModel() *Models.VpnTrafficUsage {
	return r.Model
}

func (r *VpnTrafficUsageRepository) Persist(usage Models.VpnTrafficUsage) Models.VpnTrafficUsage {
	var sql string
	var err error

	sql = "INSERT INTO vpn_traffic_usage(client_id, traffic) VALUES (?, ?)"
	_, err = r.Db.Exec(sql, usage.ClientId, usage.Traffic)
	if err != nil {
		panic(err)
	}

	return usage
}

func (r *VpnTrafficUsageRepository) BuildModel(clientId int, traffic int) Models.VpnTrafficUsage {
	return Models.VpnTrafficUsage{ClientId: clientId, Traffic: traffic}
}
