package api

import (
	"encoding/json"
	"go_cassandra/db/models"
	"go_cassandra/db/query"
	"net/http"
)

var user models.User

// Adduser endpoint to add user to the database
func Adduser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&user)
		query.AddUser(&user)

		println("just added user to database")
		json.NewEncoder(w).Encode("You just added user1 to the cassandra database")
	})
}
