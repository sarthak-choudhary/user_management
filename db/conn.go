package db

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

//Conn contains the parameters of a database connection
type Conn struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var connection Conn

//SetupDBConnection connects the server to database
func SetupDBConnection() {
	connection.cluster = gocql.NewCluster("127.0.0.1")
	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "first_app"
	connection.session, _ = connection.cluster.CreateSession()

	fmt.Println("Connected to database")
}

//ExecuteQuery executes the cql queries using session
func ExecuteQuery(query string, values ...interface{}) {
	if err := connection.session.Query(query).Bind(values...).Exec(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
