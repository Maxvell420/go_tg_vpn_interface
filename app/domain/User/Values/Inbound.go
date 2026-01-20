package values

import "GO/app/domain/User/Models"

type Inbound struct {
	Id        int
	Total     int
	CalcTotal int
	Protocol  string
	Tag       string
	Clients   []Models.VpnClient
}
