package file

import (
	"github.com/gin-gonic/gin"
	y1 "github.com/xtclalala/ScanNetWeb/api/y1/file"
)

func InitFileRouter(router *gin.RouterGroup) {
	fileRouter := router.Group("file") //.Use(middleware.LogToFile())
	{
		fileRouter.GET(":fileId", y1.Download)
		fileRouter.POST("file", y1.Upload)
	}
}
