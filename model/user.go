package model

// 用户实体类
type User struct {
	ID              uint   `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Password        string `json:"password,omitempty"`
	FollowCount     uint   `json:"follow_count,omitempty"`             //粉丝总数
	FollowerCount   uint   `json:"follower_count,omitempty"`           //关注总数
	Avatar          string `json:"avatar,omitempty"`                   //用户头像
	BackgroundImage string `json:"background_image,omitempty"`         //用户个人页顶部大图
	Signature       string `json:"signature,omitempty"`                //个人简介
	TotalFavorited  uint   `json:"total_favorited,omitempty"`          //获赞数量
	WorkCount       uint   `json:"work_count"`                         //作品数量
	FavoriteCount   uint   `json:"favorite_count,omitempty"`           //点赞数量
	FanArray        Array  `json:"type:longtext,omitempty,fan_array" ` //粉丝列表
}
