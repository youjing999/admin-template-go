package middleware

import (
	"admin-template-go/common/constant"
	"admin-template-go/common/result"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.Failed(c, int(result.ApiCode.NOAUTH), result.ApiCode.GetMessage(result.ApiCode.NOAUTH))
			c.Abort()
			return
		}
		// strings.SplitN调用将authHeader字符串（即Authorization头的值）按照空格（" "）分割成最多两个子字符串
		// 如果authHeader包含多于一个空格，它也只会在第一个空格处分割字符串。分割后的子字符串被存储在一个名为parts的切片中。
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, int(result.ApiCode.AUTHFORMATERROR), result.ApiCode.GetMessage(result.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}
		// todo 检验token
		var token = "token"
		c.Set(constant.ContextKeyUserObj, token)
		c.Next()
	}
}
