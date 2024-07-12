package controller

import (
	"admin-template-go/api/service"
	"admin-template-go/common/result"
	"github.com/gin-gonic/gin"
)

// Captcha
//@Summary 验证码接口
// @Produce json
// @Description 验证码接口
// @Success 200 {object} result.Result
// @router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := service.CaptMake()
	result.Success(c, map[string]interface{}{"idKey": id, "image": base64Image})
}
