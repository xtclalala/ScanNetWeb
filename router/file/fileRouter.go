package file

import (
	"github.com/gin-gonic/gin"
	y1 "github.com/xtclalala/ScanNetWeb/api/y1/file"
)

func InitFileRouter(router *gin.RouterGroup) {
	demoRouter := router.Group("file") //.Use(middleware.LogToFile())
	{
		demoRouter.GET("file", y1.Download)
		demoRouter.POST("file", y1.Upload)
	}
}
