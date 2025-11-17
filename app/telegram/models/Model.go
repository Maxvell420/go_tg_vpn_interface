package models

type Model interface {
	GetID() int
	GetTable() string
}
