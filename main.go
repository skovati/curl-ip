package main

import (
    "fmt"
    "net"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", mainHandler)
    r.Run(":8000")
}

func mainHandler(c *gin.Context) {
    ip, err := net.ResolveTCPAddr("tcp", c.Request.RemoteAddr)
    if err != nil {
		c.Abort()
	}
    cfIP := net.ParseIP(c.Request.Header.Get("CF-Connecting-IP"))
	if cfIP != nil {
		ip.IP = cfIP
	}
    c.String(200, fmt.Sprintln(ip.IP))
}
