package models

type Model interface {
	getID() int
	getTable() string
}
