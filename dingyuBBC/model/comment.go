package model

type Comment struct {
	ID  uint32
	CID uint32
	// 是谁评论的
	UID uint32
	// 属于哪个帖子
	PID uint32
	// 内容
	Content string
	// 是否是回复
	isReply uint32
	// 回复的哪条记录
	ReplyID uint32
	UTime   string
	CTime   string
}
