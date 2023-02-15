package repo

import (
	"errors"
	"sync"
)

type User struct {
	Id            int64      `json:"id" gorm:"id,omitempty"`
	Name          string     `json:"name" gorm:"name,omitempty"`
	FollowCount   int64      `json:"follow_count" gorm:"follow_count,omitempty"`
	FollowerCount int64      `json:"follower_count" gorm:"follower_count,omitempty"`
	IsFollow      bool       `json:"is_follow" gorm:"is_follow,omitempty"`
	User          *Login     `json:"-"`                                     //用户与账号密码之间的一对一
	Videos        []*Video   `json:"-"`                                     //用户与投稿视频的一对多
	Follows       []*User    `json:"-" gorm:"many2many:user_relations;"`    //用户之间的多对多
	FavorVideos   []*Video   `json:"-" gorm:"many2many:user_favor_videos;"` //用户与点赞视频之间的多对多
	Comments      []*Comment `json:"-"`                                     //用户与评论的一对多
}
type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDao() *UserDao {
	userOnce.Do(func() {
		userDao = new(UserDao)
	})
	return userDao
}
func (u *UserDao) QueryUserInfoById(userId int64, user *User) error {
	if user == nil {
		return errors.New("空指针错误")
	}
	Db.Where("id=?", userId).First(user)
	//Db.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(user)
	if user.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

func (u *UserDao) AddUserInfo(user *User) error {
	if user == nil {
		return errors.New("空指针错误")
	}
	return Db.Create(user).Error
}
