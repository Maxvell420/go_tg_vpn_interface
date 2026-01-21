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

func (r *VpnClientRepository) Persist(client *Models.VpnClient) {
	var sql string
	var err error

	if client.GetID() != nil {
		sql = "UPDATE vpn_clients SET enabled = ?, total = ?, remaining = ?, last_online = ?, uuid = ?, email = ?, inbound_id = ?, user_id = ?, timestamp_expire = ? WHERE id = ?"
		_, err = r.Db.Exec(sql, client.Enabled, client.Total, client.Remaining, client.LastOnline, client.Uuid, client.Email, client.InboundId, client.UserId, client.TimestampExpire, client.GetID())
		if err != nil {
			panic(err)
		}
	} else {
		sql = "INSERT INTO vpn_clients(enabled, total, remaining, last_online, uuid, email, inbound_id, user_id, timestamp_expire) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
		_, err = r.Db.Exec(sql, client.Enabled, client.Total, client.Remaining, client.LastOnline, client.Uuid, client.Email, client.InboundId, client.UserId, client.TimestampExpire)
		if err != nil {
			panic(err)
		}
	}
}

func (r *VpnClientRepository) BuildModel(id int, enabled bool, total int, remaining int, lastOnline int, uuid string, email string, inboundId int, userId int, timestampExpire int) *Models.VpnClient {
	return &Models.VpnClient{Id: &id, Enabled: &enabled, Total: &total, Remaining: &remaining, LastOnline: &lastOnline, Uuid: &uuid, Email: &email, InboundId: &inboundId, UserId: &userId, TimestampExpire: &timestampExpire}
}
