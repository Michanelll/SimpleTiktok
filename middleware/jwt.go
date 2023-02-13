package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm-learn/common"
	"gorm-learn/repo"
	"net/http"
	"time"
)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

// 获取token
func GetToken(loginUser repo.Login) (string, error) {
	claims := &Claims{
		UserId: loginUser.UserInfoId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Tiktok_Michanel",
			Subject:   "Michanel",
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString("Michanel.io")
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return "Michanel.io", nil
	})
	if token != nil {
		if key, ok := token.Claims.(*Claims); ok {
			if token.Valid {
				return key, true
			} else {
				return key, false
			}
		}
	}
	return nil, false
}

func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, common.Response{StatusCode: 401, Msg: "用户不存在"})
			c.Abort()
			return
		}
		//验证token
		token, ok := ParseToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 403,
				Msg:        "token不正确",
			})
			c.Abort()
			return
		}
		//超时
		if time.Now().Unix() > token.ExpiresAt {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 402,
				Msg:        "token过期",
			})
			c.Abort()
			return
		}
		c.Set("user_id", token.UserId)
		c.Next()
	}
}
