package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求信息
		sub := c.GetHeader("X-User-Id") // 从请求头获取用户ID
		obj := c.Request.URL.Path
		act := c.Request.Method

		// 执行授权检查
		ok, _ := e.Enforce(sub, obj, act)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
