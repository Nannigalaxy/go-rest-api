package services

import (
	"log"

	"github.com/nannigalaxy/go-rest-api/app/internal/models"
	"github.com/nannigalaxy/go-rest-api/app/internal/queries"
	"github.com/nannigalaxy/go-rest-api/app/internal/schemas"
)

func CreateUser(userData schemas.AddUserInput) (int, error) {
	userId, err := queries.InsertUser(userData)
	if err != nil {
		log.Println("Error creating user:", err)
		return -1, err
	}
	return userId, nil
}

func GetUsers() ([]models.User, error) {
	users, _ := queries.QueryUsers()
	if users == nil {
		return []models.User{}, nil
	}
	return users, nil
}

func GetUserById(userId int) (models.User, error) {
	user, err := queries.QueryUserById(userId)
	if err != nil {
		if err.Error() == "null" {
			return models.User{}, err
		}
		log.Println("Error user not found", err)
		return models.User{}, err
	}
	return user, nil
}

func RemoveUserById(userId int) error {
	err := queries.DeleteUserById(userId)
	if err != nil {
		if err.Error() == "null" {
			return err
		}
		log.Println("Error user not found", err)
		return err
	}
	return nil
}
