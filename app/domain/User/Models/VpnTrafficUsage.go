package Models

type VpnTrafficUsage struct {
	Id        int
	CreatedAt int64
	ClientId  int
	Traffic   int
}

func (v *VpnTrafficUsage) GetTable() string {
	return "vpn_traffic_usage"
}

func (v *VpnTrafficUsage) GetID() int {
	return v.Id
}
