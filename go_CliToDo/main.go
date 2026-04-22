package main

import (
	"github.com/gin-gonic/gin"
	"todo-api/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.Run(":8080")
}	