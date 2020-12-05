package services

import (
	"github.com/irisida/go-microservices/05-gin/src/domain"
	"github.com/irisida/go-microservices/05-gin/src/utils"
)

type usersService struct{}

var (
	// UsersService is exported
	UsersService usersService
)

// GetUser is a wrapper that calls the domain.GetUser
// function and passes along the provided userID to
// search. A method of the usersService type
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
