package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
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

	chView := make(chan struct{})
	chCMD := make(chan struct{})
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	go openView(chView, chCMD)

	for {
		select {
		case <-chSignal:
			chCMD <- struct{}{}
		case <-chView:
			os.Exit(0)
		}
	}

}

var defaultChromeArgs = []string{
	"--disable-background-networking",
	"--disable-background-timer-throttling",
	"--disable-backgrounding-occluded-windows",
	"--disable-breakpad",
	"--disable-client-side-phishing-detection",
	"--disable-default-apps",
	"--disable-dev-shm-usage",
	"--disable-infobars",
	"--disable-extensions",
	"--disable-features=site-per-process",
	"--disable-hang-monitor",
	"--disable-ipc-flooding-protection",
	"--disable-popup-blocking",
	"--disable-prompt-on-repost",
	"--disable-renderer-backgrounding",
	"--disable-sync",
	"--disable-translate",
	"--disable-windows10-custom-titlebar",
	"--metrics-recording-only",
	"--no-first-run",
	"--no-default-browser-check",
	"--safebrowsing-disable-auto-update",
	"--password-store=basic",
	"--use-mock-keychain",
	"--app=http://127.0.0.1:9999/static/index.html",
	"--disable-translate",
	"--disable-sync",
	"--remote-debugging-port=0",
}

func openView(chView chan struct{}, chCMD chan struct{}) {
	//chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	chromePath := lorca.LocateChrome()
	cmd := exec.Command(chromePath, defaultChromeArgs...)
	cmd.Start()

	go func() {
		<-chCMD
		cmd.Process.Kill()
	}()
	go func() {
		cmd.Wait()
		chView <- struct{}{}
	}()
}
