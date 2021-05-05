package boot

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/peter-wins/Blue-novel/global"
	"github.com/peter-wins/Blue-novel/middleware"
	"github.com/peter-wins/Blue-novel/routers"
)

func GinLaunch(){
	// 启动小说服务器,初始化gin实例
	global.Gin = gin.New()
	// 设置gin的运行模式
	gin.SetMode(gin.DebugMode)
	// 加载全局中间件
	global.Gin.Use(
		middleware.GlobalExceptionMiddleware,
		)
	// 新加的路由
	routers.InitRouter(global.Gin)

	// 服务优雅关闭和重启
	// global.Gin.Run()
	fmt.Println("Gin 服务启动成功...")
	endless.ListenAndServe(":8080", global.Gin)

	defer mysqlClose()
}

// 内置的初始化函数，服务启动时，只会执行一次
func init(){
	// 初始化全局验证器
	validateInit()
	// 初始化env全局配置
	envInit()
	// 初始化yaml配置
	yamlInit()
	// 数据库连接
	mysqlConnect()
}