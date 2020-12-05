package services

import (
	"net/http"

	"github.com/irisida/go-microservices/05-gin/src/domain"
	"github.com/irisida/go-microservices/05-gin/src/utils"
)

type itemsService struct{}

var (
	// ItemsService is exported
	ItemsService itemsService
)

// GetItem will return an item based on matching an id
// it is a method of the itemsService type
func (i *itemsService) GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "TO BE IMPLEMENTED",
		StatusCode: http.StatusInternalServerError,
	}
}
