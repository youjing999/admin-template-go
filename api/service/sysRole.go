package service

import (
	"admin-template-go/api/dao"
	"admin-template-go/api/entity"
	"admin-template-go/common/result"
	"github.com/gin-gonic/gin"
)

type ISysRoleService interface {
	CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto)
	GetSysRoleById(c *gin.Context, Id int)
	UpdateSysRole(c *gin.Context, dto entity.UpdateSysRoleDto)
	DeleteSysRoleById(c *gin.Context, dto entity.SysRoleIdDto)
	UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto)
	GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string)
	QuerySysRoleVoList(c *gin.Context)
	QueryRoleMenuIdList(c *gin.Context, Id int)
	AssignPermissions(c *gin.Context, menu entity.RoleMenu)
}

type SysRoleServiceImpl struct{}

// AssignPermissions 分配权限
func (s *SysRoleServiceImpl) AssignPermissions(c *gin.Context, menu entity.RoleMenu) {
	result.Success(c, dao.AssignPermissions(menu))
}

// QueryRoleMenuIdList 根据角色id查询菜单数据
func (s *SysRoleServiceImpl) QueryRoleMenuIdList(c *gin.Context, Id int) {
	roleMenuIdList := dao.QueryRoleMenuIdList(Id)
	var idList = make([]int, 0)
	for _, id := range roleMenuIdList {
		idList = append(idList, id.Id)
	}
	result.Success(c, idList)
}

// QuerySysRoleVoList 角色下拉列表
func (s SysRoleServiceImpl) QuerySysRoleVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysRoleVoList())
}

// GetSysRoleList 分页查询角色列表
func (s SysRoleServiceImpl) GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysRole, count := dao.GetSysRoleList(PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysRole})
	return
}

// UpdateSysRoleStatus 角色状态启用/停用
func (s SysRoleServiceImpl) UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto) {
	bool := dao.UpdateSysRoleStatus(dto)
	if !bool {
		return
	}
	result.Success(c, true)
}

// DeleteSysRoleById 根据id删除角色
func (s SysRoleServiceImpl) DeleteSysRoleById(c *gin.Context, dto entity.SysRoleIdDto) {
	dao.DeleteSysRoleById(dto)
	result.Success(c, true)
}

// UpdateSysRole 修改角色
func (s SysRoleServiceImpl) UpdateSysRole(c *gin.Context, dto entity.UpdateSysRoleDto) {
	sysRole := dao.UpdateSysRole(dto)
	result.Success(c, sysRole)
}

// GetSysRoleById 根据id查询角色
func (s SysRoleServiceImpl) GetSysRoleById(c *gin.Context, Id int) {
	sysRole := dao.GetSysRoleById(Id)
	result.Success(c, sysRole)
}

// CreateSysRole 新建角色
func (s SysRoleServiceImpl) CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto) {
	boolean := dao.CreateSysRole(dto)
	if !boolean {
		result.Failed(c, int(result.ApiCode.ROLENAMEALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.ROLENAMEALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

var sysRoleService = SysRoleServiceImpl{}

func SysRoleService() ISysRoleService {
	return &sysRoleService
}
