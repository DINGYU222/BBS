package model

type Role struct {
	ID       uint32
	RoleName string
	Users    []User
	UTime    string
	CTime    string
}
