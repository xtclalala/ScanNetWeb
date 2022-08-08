package SSH

import (
	"github.com/gin-gonic/gin"
	y1 "github.com/xtclalala/ScanNetWeb/api/y1/SSH"
)

func InitSSHRouter(router *gin.RouterGroup) {
	sshRouter := router.Group("ssh") //.Use(middleware.LogToFile())
	{
		sshRouter.GET("ssh", y1.Search)
		sshRouter.POST("ssh", y1.Create)
		sshRouter.PUT("ssh", y1.Update)
		sshRouter.DELETE("ssh", y1.Delete)
		sshRouter.POST("run", y1.Run)
		sshRouter.GET("run", y1.GetResult)
	}
}
