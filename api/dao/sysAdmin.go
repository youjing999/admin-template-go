package dao

import (
	"admin-template-go/api/entity"
	"admin-template-go/pkg/db"
)

// SysAdminDetail 用户详情
func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	db.Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}
