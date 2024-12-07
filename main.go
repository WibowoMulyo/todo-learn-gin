package main

import (
	"go-todo/config"
	"go-todo/controller"

	"github.com/gin-gonic/gin"
)

func main(){
	config.ConnectDatabase()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/todos", controller.GetTodos)
	r.GET("/todos/:id", controller.GetTodoById)
	r.POST("/todos", controller.CreateTodo)
	r.PUT("/todos/:id", controller.UpdateTodoById)
	r.DELETE("/todos/:id", controller.DeleteTodoById)

	r.Run() // listen and serve on 0.0.0.0:8080
};