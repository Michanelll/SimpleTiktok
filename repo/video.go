package repo

import "time"

type Video struct {
	Id            int64      `json:"id,omitempty"`
	UserInfoId    int64      `json:"-"`
	Author        User       `json:"author,omitempty" gorm:"-"`
	PlayUrl       string     `json:"play_url,omitempty"`
	CoverUrl      string     `json:"cover_url,omitempty"`
	FavoriteCount int64      `json:"favorite_count,omitempty"`
	CommentCount  int64      `json:"comment_count,omitempty"`
	IsFavorite    bool       `json:"is_favorite,omitempty"`
	Title         string     `json:"title,omitempty"`
	Users         []*User    `json:"-" gorm:"many2many:user_favor_videos;"`
	Comments      []*Comment `json:"-"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
}
