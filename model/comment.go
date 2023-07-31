package model

type Comment struct {
	ID         uint   `json:"id,omitempty"`
	UserID     uint   `json:"user_id,omitempty"`     //评论者id
	Content    string `json:"content,omitempty"`     //评论内容
	CreateDate string `json:"create_date,omitempty"` //评论时间
}
