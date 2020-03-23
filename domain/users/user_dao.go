package users

import (
	"fmt"
	"log"

	"github.com/nanoTitan/analytics-users-api/datasources/postgres/usersdb"
	"github.com/nanoTitan/analytics-users-api/utils/date"
	"github.com/nanoTitan/analytics-users-api/utils/errors"
)

const (
	queryInsertUser = `
		INSERT INTO users (first_name, last_name, email, date_created)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
)

var (
	usersDB = make(map[int64]*User)
)

/*
Get - The data access object function for retrieving a user given a userId
*/
func (user *User) Get() *errors.RestErr {
	// if err := usersdb.Client.Ping(); err != nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	// }

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

/*
Save - The data access object function for saving a user using it's userId
*/
func (user *User) Save() *errors.RestErr {
	user.DateCreated = date.GetNowString()
	err := usersdb.Client.QueryRow(
		queryInsertUser,
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated).Scan(&user.Id)

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	log.Println(fmt.Sprintf("adding userId %d", user.Id))
	return nil
}
