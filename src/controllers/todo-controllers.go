package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notblessy/go-ingnerd/src/config"
	"github.com/notblessy/go-ingnerd/src/models"
	"gorm.io/gorm"
)

var db *gorm.DB = config.DbConnection()

type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

func CreateTodo(c *gin.Context) {
	var data todoRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	c.JSON(http.StatusCreated, response)
}

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo

	err := db.Find(&todos)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// var response []todoResponse
	// for _, todo := range todos {
	// 	response = append(response, todoResponse{
	// 		ID:          todo.ID,
	// 		name:        todo.Name,
	// 		description: todo.Description,
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})

}
