package route

import (
	"github.com/gin-gonic/gin"
	"gorm-learn/handler/comment"
	"gorm-learn/handler/login"
	"gorm-learn/handler/user"
	"gorm-learn/handler/video"
	"gorm-learn/middleware"
)

func InitRoute() *gin.Engine {
	engine := gin.Default()
	r := engine.Group("/douyin")
	// basic apis
	r.GET("/feed/", video.FeedHandler)
	r.GET("/user/", user.GetUserInfoHandler)
	r.POST("/user/register/", middleware.Sha(), login.UserRegisterHandler)
	r.POST("/user/login/", middleware.Sha(), login.UserLoginHandler)
	r.POST("/publish/action/", video.PublishVideoHandler)
	r.GET("/publish/list/", video.GetVideoListHandler)

	// extra apis - I
	r.POST("/favorite/action/", video.PostLikeHandler)
	r.GET("/favorite/list/", video.GetLikeVideoListHandler)
	r.POST("/comment/action/", comment.PostCommentHandler)
	r.GET("/comment/list/", comment.GetCommentListHandler)

	// extra apis - II
	r.POST("/relation/action/", user.PostFollowHandler)
	r.GET("/relation/follow/list/", user.GetFollowHandler)
	r.GET("/relation/follower/list/", user.GetFollowerListHandler)

	//r.GET("/relation/friend/list/")
	//r.GET("/message/chat/")
	//r.POST("/message/action/")
	return engine
}
