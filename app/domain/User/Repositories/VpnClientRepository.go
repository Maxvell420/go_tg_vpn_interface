package Repositories

import (
	"database/sql"

	"GO/app/core/database"
	"GO/app/domain/User/Models"

	"github.com/davecgh/go-spew/spew"
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

func (r *VpnClientRepository) All() []Models.VpnClient {
	model := r.GetModel()
	table := model.GetTable()
	sql := "SELECT id,enabled,total,remaining,lastOnline,uuid,email,inbound_id,user_id,timestamp_expire FROM " + table
	rows, err := r.Db.Query(sql)
	if err != nil {
		spew.Dump(err)
	}
	defer rows.Close()
	vpnClients := []Models.VpnClient{}
	for rows.Next() {
		var vpnClient Models.VpnClient
		err = rows.Scan(&vpnClient.Id, &vpnClient.Enabled, &vpnClient.Total, &vpnClient.Remaining, &vpnClient.LastOnline, &vpnClient.Uuid, &vpnClient.Email, &vpnClient.InboundId, &vpnClient.UserId, &vpnClient.TimestampExpire)
		if err != nil {
			spew.Dump(err)
		}
		vpnClients = append(vpnClients, vpnClient)
	}
	return vpnClients
}
