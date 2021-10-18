package main

import (
	"goapi/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// data.Datatuika()
	// service.Migrate()
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
	r.GET("/api", handler.GetIiif)

	r.Run(":8060")
}
