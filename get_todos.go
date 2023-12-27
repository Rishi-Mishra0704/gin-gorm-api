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

func getTodo(c *gin.Context) {
	param := c.Param("id")
	var todo Todo
	if err := db.Where("id = ?", param).First(&todo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, todo)
	}
}
