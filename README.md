# 简要介绍
mysql数据库服务
# 使用技术
数据库：mysql
Web框架：gin
ORM框架：GORM
日志框架：Zap
# 安装部署
1.安装mysql
2.修改config下配置文件
3.切换到cmd的dbserver目录，执行go run main.go，启动服务
# 接口
http://127.0.0.1:8088/v1/addlogs
调用方式：POST
数据格式：json
