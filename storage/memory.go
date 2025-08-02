package storage

import (
	"sync"
	"todo-api/models"
)

var (
	tasks  = []models.Task{}
	nextID = 1
	mutex  sync.Mutex
)

// Get All Tasks
func GetTasks() []models.Task {
	mutex.Lock()
	defer mutex.Unlock()
	return tasks
}

// Create New Task
func CreateTask(task models.Task) models.Task {
	mutex.Lock()
	defer mutex.Unlock()
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	return task
}

// Update Task By ID
func UpdateTask(ID int, updatedTask models.Task) (bool, models.Task) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, task := range tasks {
		if task.ID == ID {
			updatedTask.ID = ID
			tasks[i] = updatedTask
			return true, updatedTask
		}
	}
	return false, models.Task{}
}

// Delete Task By ID
func DeleteTask(ID int) bool {
	mutex.Lock()
	defer mutex.Unlock()

	for i, task := range tasks {
		if task.ID == ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
