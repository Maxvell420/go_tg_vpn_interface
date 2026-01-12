package Models

type ReferalLink struct {
	Id    *int
	Hash  *string
	Tg_id *int
}

func (r *ReferalLink) GetTable() string {
	return "referal_links"
}

func (r *ReferalLink) GetID() *int {
	return r.Id
}

func (r *ReferalLink) GetHash() string {
	return *r.Hash
}

func (r *ReferalLink) GetTgId() int {
	return *r.Tg_id
}
