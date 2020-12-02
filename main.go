package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leave8080/gojson/conf"
	"github.com/leave8080/gojson/message"
)
var handler *message.Handler
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	handler = message.InitHandle(conf.Conf)
	r.POST("/data",getDass)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}