package Models

type Sale struct {
	Id          int
	CreatedAt   int64
	Value       float64
	ProductType int
	ProductId   int
	UserId      int
	Status      int
}

func (s *Sale) GetTable() string {
	return "sales"
}

func (s *Sale) GetID() int {
	return s.Id
}
