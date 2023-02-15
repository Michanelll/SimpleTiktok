package login

import (
	"github.com/gin-gonic/gin"
	"gorm-learn/common"
	"gorm-learn/proxy/loginAndRegister"
	"net/http"
)

type UserLoginResponse struct {
	common.Response
	*common.LoginResponse
}

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	raw, _ := c.Get("password")
	password, ok := raw.(string)
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{
				StatusCode: 1,
				Msg:        "密码解析错误",
			},
		})
	}
	userLoginResponse, err := loginAndRegister.QueryUserLogin(username, password)

	//用户不存在返回对应的错误
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, Msg: err.Error()},
		})
		return
	}

	//用户存在，返回相应的id和token
	c.JSON(http.StatusOK, UserLoginResponse{
		Response:      common.Response{StatusCode: 0},
		LoginResponse: userLoginResponse,
	})
}
