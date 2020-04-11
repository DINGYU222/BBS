package model

type Resoucre struct {
	ID  uint32
	RID uint32
	// 所属的板块
	SBID uint32
	// 上传的用户
	UID           uint32
	ResoucreName  string
	ResoucreURL   string
	ResoucreType  uint32
	DownloadCount uint64
}
