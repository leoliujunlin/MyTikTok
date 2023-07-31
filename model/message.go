package model

type Message struct {
	ID         uint   `json:"id,omitempty"`
	ToUserID   uint   `json:"to_user_id,omitempty"`   //消息接受者
	FromUserID uint   `json:"from_user_id,omitempty"` //消息发送者
	Content    string `json:"content,omitempty"`      //消息内容
	CreateTime string `json:"create_time,omitempty"`  //消息创建时间
}
