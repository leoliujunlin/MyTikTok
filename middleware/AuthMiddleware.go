package middleware

import (
	"TikTok/common"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 读取body中的token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取请求token
		var requestToken string
		if c.Request.Method == "GET" {
			requestToken = c.Query("token")
		} else if c.Request.Method == "POST" {
			requestToken = c.PostForm("token")
		}
		// token为空
		if requestToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 401,
				"status_msg":  "缺少token",
			})
			c.Abort()
			return
		}
		// 解析token
		token, claims, err := common.ParseToken(requestToken)
		// token非法
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 401,
				"status_msg":  "token非法",
			})
			c.Abort()
			return
		}
		// 获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.Where("id =?", userId).First(&user)
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": 401,
				"status_msg":  "token无效",
			})
			c.Abort()
			return
		}
		// 将用户信息写入上下文便于读取
		c.Set("user", user)
		c.Next()
	}
}
