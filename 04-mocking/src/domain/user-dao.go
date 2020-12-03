package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/irisida/go-microservices/04-mocking/src/utils"
)

var (
	// mock database
	users = map[int64]*User{
		123: {ID: 123, Fname: "One", Lname: "Twothree", Email: "big123@wee123.net"},
		456: {ID: 456, Fname: "Four", Lname: "Fivesix", Email: "ol456@ohaye.oi"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

// GetUser return the user or error
func (u userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("We are accessing the Database")
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "Not found",
	}
}
