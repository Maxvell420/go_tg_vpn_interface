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

func (v *VpnClient) IsUnlimited() bool {
	return v.Total == nil
}

func (v *VpnClient) IsExpired(timestamp int64) bool {
	return int64(*v.TimestampExpire) > timestamp
}
