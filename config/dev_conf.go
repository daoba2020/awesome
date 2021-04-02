package config

import (
	"awesome/pkg/logger"
	"go.uber.org/zap"
)

func initDevConf() {
	DbServer = DbServerConf{
		MySQL:        "root:root@tcp(192.168.2.88:3306)/awesome?charset=utf8&parseTime=true",
		NSQIP:        "192.168.2.88:4150",
		DBServerPort: ":8088",
	}

	logger.Leavel = zap.DebugLevel
	logger.Target = logger.File
}
