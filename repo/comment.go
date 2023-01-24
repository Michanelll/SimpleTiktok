package repo

import "time"

type Comment struct {
	Id         int64     `json:"id"`
	UserInfoId int64     `json:"-"`
	VideoId    int64     `json:"-"` //一视频对评论
	User       User      `json:"user" gorm:"-"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"-"`
	CreateDate string    `json:"create_date" gorm:"-"`
}
