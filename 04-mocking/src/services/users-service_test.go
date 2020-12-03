package services

import (
	"net/http"
	"testing"

	"github.com/irisida/go-microservices/04-mocking/src/domain"
	"github.com/irisida/go-microservices/04-mocking/src/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock         usersDaoMock
	getUserMockFunction func(int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserMockFunction(userID)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)

}
