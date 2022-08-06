package SSH

import (
	"github.com/gin-gonic/gin"
	y1 "github.com/xtclalala/ScanNetWeb/api/y1/SSH"
)

func InitSSHRouter(router *gin.RouterGroup) {
	demoRouter := router.Group("ssh") //.Use(middleware.LogToFile())
	{
		demoRouter.GET("ssh", y1.Search)
		demoRouter.POST("ssh", y1.Create)
	}
}