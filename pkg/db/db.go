package db

import (
	"awesome/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

// InitMysql 初始化MySQL
func InitMysql(dataSource string) {
	var err error
	DB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
	DB.LogMode(true)
}

// InitByTest 初始化数据库配置，仅用在单元测试
func InitByTest() {
	InitMysql(config.DbServer.MySQL)
}

