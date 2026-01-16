package Models

type Inbound struct {
	Id        *int
	Total     *int
	CalcTotal *int
	Protocol  *string
	Tag       *string
}

func (i *Inbound) GetTable() string {
	return "inbounds"
}
