package users

import (
	"encoding/json"
)

// PublicUser - a publicly facing object for user data
type PublicUser struct {
	Id          int64  `json:"user_id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser - a private facing object for user data
type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshall - marshall a slice of user objects to either PublicUser or PrivateUser based on the isPublic value
func (users Users) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

// Marshall - marshall a user object to either PublicUser or PrivateUser based on the isPublic value
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
