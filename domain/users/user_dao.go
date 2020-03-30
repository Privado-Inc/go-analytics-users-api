package users

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nanoTitan/analytics-users-api/datasources/postgres/usersdb"
	"github.com/nanoTitan/analytics-users-api/utils/errors"
	"github.com/nanoTitan/analytics-users-api/utils/pgutils"
)

const (
	queryInsertUser = `INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES ($1, $2, $3, $4, $5, $6)RETURNING id`
	queryGetUser    = `SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=$1`
	queryUpdateUser = `UPDATE users SET first_name=$2, last_name=$3, email=$4, status=$5 WHERE id=$1`
	queryDeleteUser = `DELETE FROM users WHERE id=$1`
	queryUserStatus = `SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=$1;`

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
	if scanErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); scanErr != nil {
		return pgutils.ParseError(scanErr)
	}

	log.Println(fmt.Sprintf("getting userId %d", user.Id))
	return nil
}

/*
Save - The data access object function for saving a user given it's userId
*/
func (user *User) Save() *errors.RestErr {
	saveErr := usersdb.Client.QueryRow(
		queryInsertUser,
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password).Scan(&user.Id)

	if saveErr != nil {
		return pgutils.ParseError(saveErr)
	}

	log.Println(fmt.Sprintf("adding userId %d", user.Id))
	return nil
}

/*
Update - The data access object function for updating a user
*/
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.FirstName, user.LastName, user.Email, user.Status)
	if err != nil {
		return pgutils.ParseError(err)
	}
	return nil
}

/*
Delete - The data access object function for deleting a user
*/
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return pgutils.ParseError(err)
	}
	return nil
}

// FindByStatus - query user rows given a status value
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryUserStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, pgutils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
