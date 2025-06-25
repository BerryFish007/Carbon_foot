package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取authorization header
		tokenString := context.Request.Header.Get("Authorization")
		//token为空
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token为空",
			})
			context.Abort()
			return
		}
		//错误token
		if len(tokenString) < 100 || !strings.HasPrefix(tokenString, "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "错误token",
			})
			context.Abort()
			return
		}
		//提取token的有效部分，去掉Bearer前缀
		tokenString = tokenString[7:]
		//解析token
		token, claims, err := ParseToken(tokenString)
		//非法token
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "非法token",
			})
			context.Abort()
			return
		}
		//获取claims中的信息
		uid := claims.Uid
		//写入上下文
		context.Set("curUser", uid)
		context.Next()
	}
}
