package domain

import (
	"net/http"
	"testing"

	"github.com/irisida/go-microservices/02-mvc/src/domain"
	"github.com/stretchr/testify/assert"
)

// TestGetUserNoUserFound tests for the not found
// case and expects no user to be returned and
// that an error is raised, if a user is found,
// or if no error is raised then we have a fail.
func TestGetUserNoUserFound(t *testing.T) {
	user, err := domain.GetUser(0)

	assert.Nil(t, user, "Returned a user for id: 0 when no user was expected")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Not found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

}

func TestGetUserForValidUser(t *testing.T) {
	user, err := domain.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "One", user.Fname)
	assert.EqualValues(t, "Twothree", user.Lname)
	assert.EqualValues(t, "big123@wee123.net", user.Email)
}
