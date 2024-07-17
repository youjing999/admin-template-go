package service

import (
	"admin-template-go/api/dao"
	"admin-template-go/api/entity"
	"admin-template-go/common/result"
	"admin-template-go/common/util"
	"admin-template-go/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ISysAdminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
}

type SysAdminServiceImpl struct {
}

// Login 登录
func (s SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParameter), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParameter))
		return
	}
	//验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, int(result.ApiCode.VerificationCodeHasExpired), result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}
	//校验
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.PASSWORDNOTTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}

	const status int = 2
	if sysAdmin.Status == status {
		result.Failed(c, int(result.ApiCode.STATUSISENABLE), result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}
	token, err := jwt.GenerateTokenByAdmin(sysAdmin)
	result.Success(c, map[string]interface{}{"token": token, "sysAdmin": sysAdmin})
}

var sysAdminService = SysAdminServiceImpl{}

func SysAdminService() ISysAdminService {
	return &sysAdminService
}
