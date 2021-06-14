package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/service-b/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is service B",
		})
	})
	r.Run()
}
