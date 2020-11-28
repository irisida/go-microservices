package domain

import (
	"errors"
	"fmt"
)

var (
	// mock database
	users = map[int64]*User{
		123: {ID: 123, Fname: "One", Lname: "Twothree", Email: "big123@wee123.net"},
		456: {ID: 456, Fname: "Four", Lname: "Fivesix", Email: "ol456@ohaye.oi"},
	}
)

// GetUser return the user or error
func GetUser(userID int64) (*User, error) {
	// implementation
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("user %v was not found", userID))

}
