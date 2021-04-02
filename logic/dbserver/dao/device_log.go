package dao

import (
	"awesome/logic/dbserver/model"
	"awesome/pkg/db"
	"fmt"
)

type deviceLogDao struct{}

var DeviceLogDao = new(deviceLogDao)

// Add 添加日志
func (*deviceLogDao) Add(log model.DeviceLog) (int64, error) {
	err := db.DB.Create(&log).Error
	if err != nil {
		return 0, err
	}

	return log.Id, err
}

// BatchAdd 批量添加日志
func (*deviceLogDao) BatchAdd(logs []model.DeviceLog) ([]int64, error) {
	sql := "INSERT INTO `device_log` (`project_id`,`device_id`,`level`,`content`,`create_time`,`update_time`) VALUES "
	// 循环logs数组,组合sql语句
	for key, value := range logs {
		if len(logs)-1 == key {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d,'%s','%s','%s','%s');",
				value.ProjectId, value.DeviceId, value.Level, value.Content, value.CreateTime, value.UpdateTime)
		} else {
			sql += fmt.Sprintf("(%d,%d,'%s','%s','%s','%s'),",
				value.ProjectId, value.DeviceId, value.Level, value.Content, value.CreateTime, value.UpdateTime)
		}
	}
	err := db.DB.Exec(sql).Error
	if err != nil {
		return nil, err
	}

	logIds := make([]int64, len(logs))
	for i := range logs {
		logIds[i] = logs[i].Id
	}
	return logIds, err
}


