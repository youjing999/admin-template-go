package controller

import (
	"admin-template-go/api/entity"
	"admin-template-go/api/service"
	"github.com/gin-gonic/gin"
)

// Login
// @Summary 用户登录接口
// @Produce json
// @Description 用户登录接口
// @Param data body entity.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().Login(c, dto)
}
