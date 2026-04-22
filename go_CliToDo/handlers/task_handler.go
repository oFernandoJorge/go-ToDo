package handlers

import (	
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"todo/models"
)


var tasks []models.Task
var nextID = 1

//Get /tasks

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

//Post /tasks

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	newTask.ID = nextID
	nextID++

	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

//PUT /tasks/:id

func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	for i, task := range tasks {
		if task.ID == id {
			var updatedTask models.Task

			if err := c.ShouldBindJSON(&updatedTask); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
				return
			}

			tasks[i].Title = updatedTask.Title
			tasks[i].Completed = updatedTask.Completed

			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
	}

//DELETE /tasks/:id

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)	

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Tarefa excluída com sucesso"})
			return
		}	
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
}