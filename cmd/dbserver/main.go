package main

import (
	"awesome/config"
	"awesome/logic/dbserver/model"
	"awesome/logic/dbserver/service"
	"awesome/pkg/db"
	"awesome/pkg/logger"
	"awesome/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	// 初始化日志
	logger.Init()
	// 初始化数据库
	db.InitMysql(config.DbServer.MySQL)
	// 创建表
	service.DbService.CreateTable()

	// 设置gin模式
	gin.SetMode(gin.DebugMode)

	if config.Env() == config.ENV_PROD {
		db.DB.LogMode(false)
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(Cors())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册 util.JsonTime 类型的自定义校验规则
		v.RegisterCustomTypeFunc(ValidateJSONDateType, util.JsonTime{})
	}

	// 插入数据接口
	router.POST("/v1/addlogs", handleV1AddLogs)

	router.Run(config.DbServer.DBServerPort)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, Auth")
			// 允许浏览器（客户端）可以解析的头部
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic info is: ", err)
			}
		}()

		c.Next()
	}
}

func ValidateJSONDateType(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(util.JsonTime{}) {
		timeStr := field.Interface().(util.JsonTime).String()
		// 0001-01-01 00:00:00 是go中 time.Time 类型的空值
		// 这里返回 Nil 则会被 validator 判定为空值，而无法通过 `binding:"required"` 规则
		if timeStr == "0001-01-01 00:00:00" {
			return nil
		}
		return timeStr
	}
	return nil
}

// basicAuth 鉴权
func basicAuth(c *gin.Context) error {
	return nil
}

// handleV1AddLogs 插入日志数据接口处理函数
func handleV1AddLogs(c *gin.Context) {
	// 这里需要实现一下鉴权
	err := basicAuth(c)
	if err != nil {
		c.JSON(http.StatusOK, Response{Code: http.StatusUnauthorized, Message: err.Error()})
		return
	}

	// 处理请求数据
	var logs []model.DeviceLog
	err = c.BindJSON(&logs)
	if err != nil {
		c.JSON(http.StatusOK, Response{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if len(logs) == 0 {
		c.JSON(http.StatusOK, Response{Code: http.StatusBadRequest, Message: "no data passed"})
		return
	}

	err = service.DbService.AddLogs(logs)
	if err != nil {
		c.JSON(http.StatusOK, Response{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "success", Data: "success"})
}
