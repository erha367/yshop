package controller

import (
	"github.com/gin-gonic/gin"
	"yshop/model"
	"yshop/utils"
)

func Ping(c *gin.Context) {
	utils.Info(c, "Received request: %s %s", c.Request.Method, c.Request.URL.Path)
	var cates []model.ProductCate
	utils.SQLite.WithContext(c).Where("parent_id = ?", 0).Find(&cates)
	c.JSON(200, gin.H{
		"message": cates,
	})
}
