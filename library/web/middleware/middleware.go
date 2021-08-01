package middleware

import (
	"go-app/library/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, "no route")
}

func NoMethod(c *gin.Context) {
	c.JSON(http.StatusNotFound, "no method")
}

func AccessLog(c *gin.Context) {
	start := time.Now()
	c.Next()
	end := time.Now()

	statusCode := c.Writer.Status()
	latency := end.Sub(start)
	clientIP := c.ClientIP()
	method := c.Request.Method
	path := c.Request.URL.Path
	if raw := c.Request.URL.RawQuery; raw != "" {
		path += "?" + raw
	}
	log.Accessf("[GIN] %s | %3d | %13v | %15s | %7s | %s",
		end.Format("2006/01/02 - 15:04:05.999"),
		statusCode,
		latency,
		clientIP,
		method,
		path,
	)
}
