package models

type Comment struct {
	Id       uint64
	Comment  string
	ParentId *uint64
	//Type uint8
	//ObjectId uint
	//Status uint8
	//UserId uint
}
