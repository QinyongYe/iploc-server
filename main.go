package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slene/iploc"
	"path/filepath"
	"fmt"
)

func init() {
	// replace iplocFilePath to your iploc.dat path
	iplocFilePath, _ := filepath.Abs("iploc.dat")

	// simple set a true param can preload all ipinfo
	// need allocate more memory > 30M
	// and speed can grow up about 40 percent than not preload
	iploc.IpLocInit(iplocFilePath, true)

	// read iploc.dat into memory, not preload
	// iploc.IpLocInit(iplocFilePath)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("client ip: ", c.ClientIP())
		ipinfo, err := iploc.GetIpInfo(c.ClientIP())
		if err != nil {
			c.JSON(500, err)
		} else {
			c.JSON(200, ipinfo)
		}
	})
	r.Run(":8081") // listen and server on 0.0.0.0:8080
}
