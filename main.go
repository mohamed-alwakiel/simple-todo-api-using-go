package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-api/handlers"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/tasks", handlers.GetTasks)
	router.POST("/tasks", handlers.CreateTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)

	err := router.Run(":8080")
	if err != nil {
		fmt.Print(err)
		return
	} // شغل السيرفر على بورت 8080
}
