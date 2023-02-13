package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
)

func sha(str string) string {
	w := sha1.New()
	io.WriteString(w, str)           //将str写入到w中
	bw := w.Sum(nil)                 //w.Sum(nil)将w的hash转成[]byte格式
	shaStr := hex.EncodeToString(bw) //将 bw 转成字符串
	return shaStr
}
func Sha() gin.HandlerFunc {
	return func(context *gin.Context) {
		password := context.Query("password")
		if password == "" {
			password = context.PostForm("password")
		}
		context.Set("password", sha(password))
		context.Next()
	}
}
