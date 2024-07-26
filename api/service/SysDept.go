package service

import (
	"admin-template-go/api/dao"
	"admin-template-go/api/entity"
	"admin-template-go/common/result"
	"github.com/gin-gonic/gin"
)

type ISysDeptService interface {
	GetSysDeptList(c *gin.Context, DeptName, DeptStatus string)
	CreateSysDept(c *gin.Context, sysDept entity.SysDept)
	GetSysDeptById(c *gin.Context, Id int)
	UpdateSysDept(c *gin.Context, dept entity.SysDept)
	DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto)
	QuerySysDeptVoList(c *gin.Context)
}

type SysDeptServiceImpl struct{}

// QuerySysDeptVoList 部门下拉列表
func (s SysDeptServiceImpl) QuerySysDeptVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysDeptVoList())
}

// DeleteSysDeptById 删除部门
func (s SysDeptServiceImpl) DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto) {
	boolean := dao.DeleteSysDeptById(dto)
	if !boolean {
		result.Failed(c, int(result.ApiCode.DEPTISDISTRIBUTE), result.ApiCode.GetMessage(result.ApiCode.DEPTISDISTRIBUTE))
		return
	}
	result.Success(c, true)
}

// UpdateSysDept 修改部门
func (s SysDeptServiceImpl) UpdateSysDept(c *gin.Context, dept entity.SysDept) {
	sysDept := dao.UpdateSysDept(dept)
	result.Success(c, sysDept)
}

// GetSysDeptById 根据id查询部门
func (s SysDeptServiceImpl) GetSysDeptById(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysDeptById(Id))
}

// CreateSysDept 新增部门
func (s SysDeptServiceImpl) CreateSysDept(c *gin.Context, sysDept entity.SysDept) {
	bool := dao.CreateSysDept(sysDept)
	if !bool {
		result.Failed(c, int(result.ApiCode.DEPTISEXIST), result.ApiCode.GetMessage(result.ApiCode.DEPTISEXIST))
		return
	}
	result.Success(c, true)
}

// GetSysDeptList 查询部门列表
func (s SysDeptServiceImpl) GetSysDeptList(c *gin.Context, DeptName, DeptStatus string) {
	result.Success(c, dao.GetSysDeptList(DeptName, DeptStatus))
}

var sysDeptService = SysDeptServiceImpl{}

func SysDeptService() ISysDeptService {
	return &sysDeptService
}
