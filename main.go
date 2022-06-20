package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

//go:embed web/dist/*
var FS embed.FS

func main() {

	go func() {
		staticFile, _ := fs.Sub(FS, "web/dist")
		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		router.StaticFS("/static", http.FS(staticFile))
		router.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path
			if strings.HasPrefix(path, "/static/") {
				reader, err := staticFile.Open("index.html")
				if err != nil {
					c.Status(http.StatusNotFound)
					return
				}
				defer reader.Close()
				stat, err := reader.Stat()
				if err != nil {
					c.Status(http.StatusNotFound)
					return
				}
				c.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil)
			} else {
				c.Status(http.StatusNotFound)
			}
		})
		router.Run(":9999")
	}()

	chromePath := "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:9999/static/index.html  --disable-translate --disable-sync")
	cmd.Start()

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	select {
	case <-chSignal:
		cmd.Process.Kill()
	}
}
