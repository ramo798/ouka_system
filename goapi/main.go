package main

import (
	"goapi/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "living",
		})
	})
	r.GET("/all", handler.Getall)
	r.GET("/serch", handler.Getserch)

	r.Run()
}
