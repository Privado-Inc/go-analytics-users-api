package users

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nanoTitan/analytics-users-api/datasources/postgres/usersdb"
	"github.com/nanoTitan/analytics-users-api/utils/date"
	"github.com/nanoTitan/analytics-users-api/utils/errors"
	"github.com/nanoTitan/analytics-users-api/utils/pgutils"
)

const (
	queryInsertUser = `
		INSERT INTO users (first_name, last_name, email, date_created)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	queryGetUser = `
		SELECT id, first_name, last_name, email, date_created 
		FROM users
		WHERE id=$1
	`

	indexUniqueEmail = "users_email_key"
	errorNoRows      = "no rows in result set"
)

/*
Get - The data access object function for retrieving a user given a userId
*/
func (user *User) Get() *errors.RestErr {
	stmt, prepErr := usersdb.Client.Prepare(queryGetUser)
	if prepErr != nil {
		return errors.NewInternalServerError(prepErr.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if scanErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); scanErr != nil {
		return pgutils.ParseError(scanErr)
	}

	log.Println(fmt.Sprintf("getting userId %d", user.Id))
	return nil
}

/*
Save - The data access object function for saving a user given it's userId
*/
func (user *User) Save() *errors.RestErr {
	user.DateCreated = date.GetNowString()
	saveErr := usersdb.Client.QueryRow(
		queryInsertUser,
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated).Scan(&user.Id)

	if saveErr != nil {
		return pgutils.ParseError(saveErr)
	}

	log.Println(fmt.Sprintf("adding userId %d", user.Id))
	return nil
}
