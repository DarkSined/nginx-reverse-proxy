package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/service-a/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is service A",
		})
	})
	r.Run()
}
