package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/xtclalala/ScanNetWeb/internal/net"
	service "github.com/xtclalala/ScanNetWeb/services/ws"
	"log"
)

func WsConnect(c *gin.Context) {
	log.Println("websocket链接")
	service.WsHandler(c.Writer, c.Request)
}

func Delect(c *gin.Context) {
	if err := service.DeleteClient(); err != nil {
		net.FailWithMessage("未找到该客户端", c)
		return
	}
	net.Ok(c)
}
