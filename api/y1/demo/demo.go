package y1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xtclalala/ScanNetWeb/internal/net"
	"github.com/xtclalala/ScanNetWeb/services/demo"
)

func DoAnyThing(c *gin.Context) {
	fmt.Println("do any thing with ctl")
	demo.DoAnything()
	net.Ok(c)
	return
}
