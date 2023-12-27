package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func createTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)

	if err := db.Create(&todo).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, todo)
	}
}
