package models

type Comment struct {
	Id       uint64
	Comment  string
	ParentId *uint64
}
