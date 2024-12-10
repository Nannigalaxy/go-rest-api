package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/nannigalaxy/go-rest-api/app/connections/database"
	"github.com/nannigalaxy/go-rest-api/app/internal/models"
	"github.com/nannigalaxy/go-rest-api/app/internal/schemas"
)

var db = database.DBConnection
var ctx = database.DBContext

func InsertUser(userData schemas.AddUserInput) (int, error) {
	query := "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id;"
	var user_id int
	err := db.QueryRow(ctx, query, userData.Username, userData.Email).Scan(&user_id)
	if err != nil {
		log.Printf("Error querying user: %v", err)
		return -1, err
	}
	return user_id, nil
}

func QueryUsers() ([]models.User, error) {
	query := "SELECT id, username, email FROM users"
	rows, err := db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			log.Printf("Something wrong while reading %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func QueryUserById(userId int) (models.User, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	row := db.QueryRow(ctx, query, userId)
	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No user found with ID %d", userId)
			return models.User{}, fmt.Errorf("null")
		}
		log.Printf("Error querying users: %v", err)
		return models.User{}, err
	}
	return user, nil
}

func DeleteUserById(userId int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := db.Exec(ctx, query, userId)
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No user found with ID %d", userId)
		return fmt.Errorf("null")
	}
	return err
}
