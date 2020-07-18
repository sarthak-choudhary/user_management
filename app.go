package main

import (
	"go_cassandra/db"
)

//InitializeApp - initialises the server
func InitializeApp() {
	db.SetupDBConnection()
}
