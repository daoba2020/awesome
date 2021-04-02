package model

import (
	"awesome/pkg/util"
)

// DeviceLog 日志
type DeviceLog struct {
	Id         int64          `gorm:"not null;type:bigint(20) primary key auto_increment;comment:'自增主键'"`
	ProjectId  int64          `gorm:"not null;type:bigint(20);index:idx_type_id;comment:'项目ID'"`
	DeviceId   int64          `gorm:"not null;type:bigint(20);index:idx_type_id;comment:'设备ID'"`
	Level      string         `gorm:"not null;type:varchar(16);comment:'日志级别'"`
	Content    string         `gorm:"not null;type:varchar(256);comment:'内容'"`
	CreateTime *util.JsonTime `gorm:"type:datetime default current_timestamp;comment:'创建时间'"`
	UpdateTime *util.JsonTime `gorm:"type:datetime default current_timestamp on update current_timestamp;comment:'更新时间'"`
}
