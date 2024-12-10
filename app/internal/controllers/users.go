package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"

	"github.com/nannigalaxy/go-rest-api/app/internal/schemas"
	"github.com/nannigalaxy/go-rest-api/app/internal/services"
)

// Controller functions

func CreateUser(c *gin.Context) {
	var newUser schemas.AddUserInput
	// validate input with schema
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, err := services.CreateUser(newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.String(http.StatusOK, strconv.Itoa(user_id))

}

func GetAllUsers(c *gin.Context) {
	users, _ := services.GetUsers()
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var userId_str string = c.Param("user_id")
	userId, _ := strconv.Atoi(userId_str)
	users, err := services.GetUserById(userId)
	if err != nil {
		if err.Error() == "null" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		log.Printf("Error fetching user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func RemoveUser(c *gin.Context) {
	var userId_str string = c.Param("user_id")
	userId, _ := strconv.Atoi(userId_str)
	err := services.RemoveUserById(userId)
	if err != nil {
		if err.Error() == "null" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		log.Printf("Error fetching user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.String(http.StatusOK, userId_str)
}

// User routes mapping
func UserRouter(router *gin.RouterGroup) {
	user := router.Group("/users")
	{
		user.GET("", GetAllUsers)
		user.POST("", CreateUser)
		user.GET(":user_id", GetUser)
		user.DELETE(":user_id", RemoveUser)
	}

}
