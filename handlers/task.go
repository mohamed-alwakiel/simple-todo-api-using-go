package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-api/models"
	"todo-api/storage"
)

// GET /tasks
func GetTasks(g *gin.Context) {
	tasks := storage.GetTasks()
	g.JSON(http.StatusOK, tasks)
}

// POST /tasks
func CreateTask(g *gin.Context) {
	var task models.Task
	if err := g.ShouldBindJSON(&task); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask := storage.CreateTask(task)
	g.JSON(http.StatusCreated, createdTask)
}

// PUT /tasks/:id
func UpdateTask(g *gin.Context) {
	ID, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := g.ShouldBindJSON(&task); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, updatedTask := storage.UpdateTask(ID, task)
	if !ok {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}

	g.JSON(http.StatusOK, updatedTask)
}

// DELETE /tasks/:id
func DeleteTask(g *gin.Context) {
	ID, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	ok := storage.DeleteTask(ID)
	if !ok {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
