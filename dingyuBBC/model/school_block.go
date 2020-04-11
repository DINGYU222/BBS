package model

type SchoolBlock struct {
	ID          uint32
	SchoolName  string
	BlockName   string
	PeopleCount uint64
	Users       []User
	// 这个板块下的资源
	Resoucres []Resoucre
	// 这个板块下的帖子
	Posts         []Post
	ResourceCount uint64
	UTime         string
	CTime         string
}
