package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func updateTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&todo)
	db.Save(&todo)
	c.JSON(200, todo)
}

func deleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	d := db.Where("id = ?", id).Delete(&todo)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
