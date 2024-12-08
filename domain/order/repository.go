package order

import (
	"applicationDesignTest/entity"
	"errors"
)

var (
	// ErrFailedToAddOrder is returned when order could not be added to the repository.
	ErrFailedToAddOrder = errors.New("failed to add order to the repository")
)

type OrderRepository interface {
	Add(*entity.Order) error
	Get(string) (entity.Order, bool)
}
