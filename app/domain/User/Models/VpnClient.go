package Models

type VpnClient struct {
	Id              *int
	Enabled         *bool
	Total           *int
	Remaining       *int
	LastOnline      *int
	Uuid            *string
	Email           *string
	InboundId       *int
	UserId          *int
	TimestampExpire *int
}

func (v *VpnClient) GetTable() string {
	return "vpn_clients"
}

func (v *VpnClient) GetID() *int {
	return v.Id
}
