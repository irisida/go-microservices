package services

import (
	"net/http"
	"testing"

	"github.com/irisida/go-microservices/04-mocking/src/domain"
	"github.com/irisida/go-microservices/04-mocking/src/utils"
	"github.com/stretchr/testify/assert"
)

type usersDaoMock struct{}

var (
	userDaoMock         usersDaoMock
	getUserMockFunction func(int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserMockFunction(userID)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserMockFunction = func(i int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 was not found",
		}
	}

	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserMockFunction = func(i int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID:    123,
			Fname: "Mocky",
			Lname: "McMockface",
			Email: "mockytalkie@mockmock.rock",
		}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Mocky", user.Fname)
	assert.EqualValues(t, "McMockface", user.Lname)

}
