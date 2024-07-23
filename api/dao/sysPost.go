package dao

import (
	"admin-template-go/api/entity"
	"admin-template-go/common/util"
	. "admin-template-go/pkg/db"
	"time"
)

// GetSysPostByCode 根据编码查询
func GetSysPostByCode(postCode string) (sysPost entity.SysPost) {
	Db.Where("post_code = ?", postCode).First(&sysPost)
	return sysPost
}

// GetSysPostByName 根据名称查询
func GetSysPostByName(postName string) (sysPost entity.SysPost) {
	Db.Where("post_name = ?", postName).First(&sysPost)
	return sysPost
}

// CreateSysPost 新增岗位
func CreateSysPost(sysPost entity.SysPost) bool {
	sysPostByCode := GetSysPostByCode(sysPost.PostCode)
	if sysPostByCode.ID > 0 {
		return false
	}
	sysPostByName := GetSysPostByName(sysPost.PostName)
	if sysPostByName.ID > 0 {
		return false
	}
	addSysPost := entity.SysPost{
		PostCode:   sysPost.PostCode,
		PostName:   sysPost.PostName,
		PostStatus: sysPost.PostStatus,
		CreateTime: util.HTime{Time: time.Now()},
		Remark:     sysPost.Remark,
	}
	tx := Db.Save(&addSysPost)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetSysPostList 分页查询岗位列表
func GetSysPostList(PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) (sysPost []entity.SysPost, count int64) {
	curDb := Db.Table("sys_post")
	if PostName != "" {
		curDb = curDb.Where("post_name = ?", PostName)
	}
	if PostStatus != "" {
		curDb = curDb.Where("post_status = ?", PostStatus)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("`create_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysPost)
	return sysPost, count
}

// GetSysPostById 根据id查询岗位
func GetSysPostById(Id int) (sysPost entity.SysPost) {
	Db.First(&sysPost, Id)
	return sysPost
}

// UpdateSysPost 修改岗位
func UpdateSysPost(post entity.SysPost) (sysPost entity.SysPost) {
	Db.First(&sysPost, post.ID)
	sysPost.PostName = post.PostName
	sysPost.PostCode = post.PostCode
	sysPost.PostStatus = post.PostStatus
	if post.Remark != "" {
		sysPost.Remark = post.Remark
	}
	Db.Save(&sysPost)
	return sysPost
}

// DeleteSysPostById 根据id删除岗位
func DeleteSysPostById(dto entity.SysPostIdDto) {
	Db.Delete(&entity.SysPost{}, dto.Id)
}

// BatchDeleteSysPost 批量删除岗位
func BatchDeleteSysPost(dto entity.DelSysPostDto) {
	Db.Where("id in (?)", dto.Ids).Delete(&entity.SysPost{})
}

// UpdateSysPostStatus 修改状态
func UpdateSysPostStatus(dto entity.UpdateSysPostStatusDto) {
	var sysPost entity.SysPost
	Db.First(&sysPost, dto.Id)
	sysPost.PostStatus = dto.PostStatus
	Db.Save(&sysPost)
}

// QuerySysPostVoList 岗位下拉列表
func QuerySysPostVoList() (sysPostVo []entity.SysPostVo) {
	Db.Table("sys_post").Select("id, post_name").Scan(&sysPostVo)
	return sysPostVo
}
