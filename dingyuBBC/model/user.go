package model

type User struct {
	ID       uint32
	UID      uint32
	Username string
	Password string
	Phone    string
	Roles    []Role
	// 用户上传的资源
	Resoucres []Resoucre
	UTime     string
	CTime     string
}
