package handler

import (
	"goapi/service"

	"github.com/gin-gonic/gin"
)

func Getall(c *gin.Context) {
	all := service.GetAllColumns()

	c.JSON(200, all)
}
