package model

type Post struct {
	ID  uint32
	PID uint32
	// 发帖者
	UID uint32
	// 属于哪个板块
	SBID    uint32
	Title   string
	Content string
	// 评论
	Comments     []Comment
	CommentCount uint64
	DiggCount    uint64
	ShareCount   uint64
}
