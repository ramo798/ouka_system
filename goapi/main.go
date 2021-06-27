package main

import (
	"goapi/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "living",
		})
	})
	r.GET("/all", handler.Getall)

	r.Run()
}
