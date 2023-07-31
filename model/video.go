package model

import "time"

type Video struct {
	ID             uint      `json:"id"`
	AuthorID       uint      `json:"Author_id"`      //作者ID
	PlayURL        string    `json:"play_url"`       //视频的地址
	CoverURL       string    `json:"cover_url"`      //封面图地址
	FavoriteCount  uint      `json:"favorite_count"` //点赞数
	CommentCount   uint      `json:"comment_count"`  //评论数
	Title          string    `json:"title"`          //标题
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	LikeArray      Array     `json:"type:longtext,like_array"`       //粉丝列表
	CommentIDArray Array     `json:"type:longtext,comment_id_array"` //评论列表
}
