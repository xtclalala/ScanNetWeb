package demo

import (
	"github.com/gin-gonic/gin"
	y1 "github.com/xtclalala/ScanNetWeb/api/y1/demo"
)

func InitDemoRouter(router *gin.RouterGroup) {
	demoRouter := router.Group("demo") //.Use(middleware.LogToFile())
	{
		demoRouter.GET("demo", y1.DoAnyThing)
	}
}
