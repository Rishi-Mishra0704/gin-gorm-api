package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getTodos(c *gin.Context) {
	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, todos)
	}
}
