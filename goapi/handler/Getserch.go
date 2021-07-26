package handler

import (
	"goapi/service"

	"github.com/gin-gonic/gin"
)

func Getserch(c *gin.Context) {
	text := c.Query("text")
	subject := c.Query("subject")
	sort := c.DefaultQuery("sort", "ASC")

	res := service.GetQuery(text, subject, sort)

	c.JSON(200, res)
}
