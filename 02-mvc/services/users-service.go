package services

import "github.com/irisida/go-microservices/02-mvc/domain"

// GetUser is a wrapper that calls the
// domain.GetUser function and passes
// along the provided userID to search.
func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
