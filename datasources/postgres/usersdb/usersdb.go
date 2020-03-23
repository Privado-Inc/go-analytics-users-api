package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	postgresUsersHost         = "postgres_users_host"
	postgresUsersPost         = "postgres_users_port"
	postgresUsersUser         = "postgres_users_user"
	postgresUsersHostPassword = "postgres_users_password"
	postgresUsersHostDbname   = "postgres_users_dbname"
)

var (
	// Client - The postgres user db client
	Client *sql.DB

	host     = os.Getenv(postgresUsersHost)
	port     = os.Getenv(postgresUsersPost)
	user     = os.Getenv(postgresUsersUser)
	password = os.Getenv(postgresUsersHostPassword)
	dbname   = os.Getenv(postgresUsersHostDbname)
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println(fmt.Sprintf("attempting to connect to %s", psqlInfo))

	// Calling open here just evaluates the arguments
	var err error
	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Must ping the db to actually open the connection
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("datbase successfully configured")
}
