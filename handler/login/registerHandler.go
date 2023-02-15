package login

import (
	"github.com/gin-gonic/gin"
	"gorm-learn/common"
	"gorm-learn/proxy/loginAndRegister"
	"net/http"
)

type RegisterResponse struct {
	common.Response
	*common.LoginResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	rawVal, _ := c.Get("password")
	password, ok := rawVal.(string)
	if !ok {
		c.JSON(http.StatusOK, RegisterResponse{
			Response: common.Response{
				StatusCode: 1,
				Msg:        "密码解析出错",
			},
		})
		return
	}
	registerResponse, err := loginAndRegister.PostUserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, RegisterResponse{
			Response: common.Response{
				StatusCode: 1,
				Msg:        err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{
		Response:      common.Response{StatusCode: 0},
		LoginResponse: registerResponse,
	})
}
