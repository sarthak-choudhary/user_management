package query

import (
	"go_cassandra/db"
	"go_cassandra/db/models"
)

// AddUser adds a new user to the database
func AddUser(user *models.User) {
	query := `INSERT INTO users(first_name, last_name, email) VALUES (?, ?, ?)`
	db.ExecuteQuery(query, user.FirstName, user.LastName, user.Email)
}
