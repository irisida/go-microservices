package services

import (
	"github.com/irisida/go-microservices/03-testing/src/domain"
	"github.com/irisida/go-microservices/03-testing/src/utils"
)

// GetUser is a wrapper that calls the
// domain.GetUser function and passes
// along the provided userID to search.
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
