package service

import (
	"awesome/logic/dbserver/dao"
	"awesome/logic/dbserver/model"
	"awesome/pkg/db"
	"awesome/pkg/logger"
	"go.uber.org/zap"
)

type dbService struct {}

var DbService = new(dbService)

// CreateTable 根据model定义创建mysql表
func (*dbService) CreateTable() error {
	if !db.DB.HasTable(&model.DeviceLog{}) {
		db.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT COLLATE=utf8mb4_bin COMMENT='日志'").AutoMigrate(&model.DeviceLog{})
		if !db.DB.HasTable(&model.DeviceLog{}) {
			panic("创建表device_log失败")
		}
	}

	return nil
}

// AddLog 插入单条日志
func (*dbService) AddLog(log model.DeviceLog) error {
	_, err := dao.DeviceLogDao.Add(log)
	if err != nil {
		logger.Logger.Error("插入单条日志失败", zap.Any("error", err))
		return err
	}

	return nil
}

// AddLogs 插入多条日志
func (*dbService) AddLogs(logs []model.DeviceLog) error {
	_, err := dao.DeviceLogDao.BatchAdd(logs)
	if err != nil {
		logger.Logger.Error("插入多条日志失败", zap.Any("error", err))
		return err
	}

	return nil
}
