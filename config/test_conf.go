package config

import (
	"awesome/pkg/logger"
	"go.uber.org/zap"
)

func initTestConf() {
	DbServer = DbServerConf{
		MySQL:        "root:root@tcp(192.168.2.88:3306)/gim?charset=utf8&parseTime=true",
		NSQIP:        "192.168.2.88:4150",
		DBServerPort: ":8088",
	}

	logger.Leavel = zap.DebugLevel
	logger.Target = logger.File
}
