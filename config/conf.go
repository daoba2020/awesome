package config

import (
	"os"
)

const (
	IM_ENV   = "IM_ENV"
	ENV_DEV  = "dev"
	ENV_TEST = "test"
	ENV_PROD = "prod"
)

var (
	DbServer DbServerConf // DB服务相关配置
)

var environment string // 系统环境类型

// DbServer配置
type DbServerConf struct {
	MySQL        string // mysql连接字符串
	NSQIP        string // nsq服务地址
	DBServerPort string // DB服务端口
}

func init() {
	env := os.Getenv(IM_ENV)
	if env == "" {
		env = ENV_TEST
	}

	environment = env

	switch env {
	case ENV_DEV:
		initDevConf()
	case ENV_TEST:
		initTestConf()
	case ENV_PROD:
		initProdConf()
	default:
		initTestConf()
	}
}

// Env 系统环境类型
func Env() string {
	return environment
}
