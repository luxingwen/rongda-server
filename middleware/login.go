package middleware

import (
	"net/http"
	"sgin/pkg/app"
	"sgin/pkg/utils"
)

// 登录中间件
func LoginCheck() app.HandlerFunc {
	return func(c *app.Context) {

		// 获取token
		token := c.GetHeader("X-Token")
		if token == "" {
			token = c.GetHeader("Wx-Token")
		}

		if token == "" {
			c.JSONError(http.StatusUnauthorized, "未登录")
			c.Abort()
			return
		}

		// 根据token获取用户信息
		userId, wxUserId, err := utils.ParseTokenGetUserID(token)
		if err != nil {
			c.JSONError(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// 将用户信息放入上下文
		c.Set("user_id", userId)
		c.Set("wx_user_id", wxUserId)
	}
}
