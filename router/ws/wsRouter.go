package ws

import (
	"github.com/gin-gonic/gin"
	y1 "github.com/xtclalala/ScanNetWeb/api/y1/ws"
)

func InitWsRouter(router *gin.RouterGroup) {
	wsRouter := router.Group("ws") //.Use(middleware.LogToFile())
	{
		wsRouter.GET("", y1.WsConnect)
		wsRouter.DELETE("", y1.Delect)

	}
}
