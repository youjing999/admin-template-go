package dao

import "admin-template-go/api/entity"
import . "admin-template-go/pkg/db"

// CreateSysOperationLog 新增操作日志
func CreateSysOperationLog(log entity.SysOperationLog) {
	Db.Create(&log)
}

// GetSysOperationLogList 分页查询操作日志列表
func GetSysOperationLogList(Username, BeginTime, EndTime string, PageSize, PageNum int) (sysOperationLog []entity.SysOperationLog, count int64) {
	curDb := Db.Table("sys_operation_log")
	if Username != "" {
		curDb = curDb.Where("username =?", Username)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("`create_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysOperationLog)
	return sysOperationLog, count
}

// DeleteSysOperationLogById 根据id删除操作日志
func DeleteSysOperationLogById(dto entity.SysOperationLogIdDto) {
	Db.Delete(&entity.SysOperationLog{}, dto)
}

// BatchDeleteSysOperationLog 批量删除批量操作日志
func BatchDeleteSysOperationLog(dto entity.BatchDeleteSysOperationLogDto) {
	Db.Where("id in (?)", dto.Ids).Delete(&entity.SysOperationLog{})
}

// CleanSysOperationLog 清空操作日志
func CleanSysOperationLog() {
	Db.Exec("truncate table sys_operation_log")
}
