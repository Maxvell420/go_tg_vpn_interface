package Repositories

import (
	"database/sql"

	"GO/app/core/database"
	"GO/app/domain/User/Models"
)

type VpnClientRepository struct {
	database.Repository
	Model *Models.VpnClient
	Db    *sql.DB
}

func (r *VpnClientRepository) GetModel() *Models.VpnClient {
	return r.Model
}

func (r *VpnClientRepository) GetDB() *sql.DB {
	return r.Db
}

func (r *VpnClientRepository) AllUuidMap() (map[string]Models.VpnClient, error) {
	model := r.GetModel()
	table := model.GetTable()
	uuidMap := make(map[string]Models.VpnClient)
	sql := "SELECT id,enabled,total,remaining,lastOnline,uuid,email,inbound_id,user_id,timestamp_expire FROM " + table
	rows, err := r.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var vpnClient Models.VpnClient
		err = rows.Scan(&vpnClient.Id, &vpnClient.Enabled, &vpnClient.Total, &vpnClient.Remaining, &vpnClient.LastOnline, &vpnClient.Uuid, &vpnClient.Email, &vpnClient.InboundId, &vpnClient.UserId, &vpnClient.TimestampExpire)
		if err != nil {
			return nil, err
		}
		uuidMap[*vpnClient.Uuid] = vpnClient
	}
	return uuidMap, nil
}
