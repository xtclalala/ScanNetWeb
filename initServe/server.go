package initServe

import (
	"github.com/gin-gonic/gin"
	"github.com/xtclalala/ScanNetWeb/middleware"
	"github.com/xtclalala/ScanNetWeb/router/SSH"
	"github.com/xtclalala/ScanNetWeb/router/demo"
	"github.com/xtclalala/ScanNetWeb/router/file"
)

func InitApi(router *gin.Engine) {
	// 公共路由
	router.Use(middleware.LogToFile())
	publicGroup := router.Group("")
	//var userApi = v1.ApiGroupApp.SysApiGroup.UserApi
	{
		//publicGroup.POST("login", )
		demo.InitDemoRouter(publicGroup)

	}
	// 私有路由
	router.Group("")
	//router.Use(xxxxx)
	privateGroup := router.Group("")
	SSH.InitSSHRouter(privateGroup) // linux ssh
	file.InitFileRouter(privateGroup)
	//sysRouter := Router.AppRouter.System
	//sysRouter.InitUserRouter(privateGroup)     // 用户路由
	//sysRouter.InitPerRouter(privateGroup)      // 按钮路由
	//sysRouter.InitMenuRouter(privateGroup)     // 菜单路由
	//sysRouter.InitRoleRouter(privateGroup)     // 角色路由
	//sysRouter.InitOrganizeRouter(privateGroup) // 组织路由
}
