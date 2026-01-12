package Models

type ReferalUser struct {
	Id          *int
	Tg_id       *int
	Owner_tg_id *int
}

func (r *ReferalUser) GetTable() string {
	return "referal_users"
}

func (r *ReferalUser) GetID() *int {
	return r.Id
}

func (r *ReferalUser) GetTgId() *int {
	return r.Tg_id
}

func (r *ReferalUser) GetOwnerTgId() *int {
	return r.Owner_tg_id
}
