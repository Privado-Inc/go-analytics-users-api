package services

import (
	"github.com/nanoTitan/analytics-users-api/domain/users"
	"github.com/nanoTitan/analytics-users-api/utils/cryptoutils"
	"github.com/nanoTitan/analytics-users-api/utils/date"
	"github.com/nanoTitan/analytics-users-api/utils/errors"
)

// GetUser - User service to get a user object through the data access object
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateUser - User service to create a user object through the data access object
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date.GetNowDbFormat()
	user.Status = users.StatusActive
	user.Password = cryptoutils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser - User service to update a user object through the data access object
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err = current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser - delete a user object given a user ID
func DeleteUser(usersID int64) *errors.RestErr {
	user := &users.User{Id: usersID}
	return user.Delete()
}

// Search - get rows of user objects given a status string
func Search(status string) (users.Users, *errors.RestErr) {
	user := &users.User{}
	return user.FindByStatus(status)
}
