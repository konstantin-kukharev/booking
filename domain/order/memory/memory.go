package memory

import (
	"applicationDesignTest/domain/order"
	"applicationDesignTest/entity"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// MemoryRepository fulfills the OrderRepository interface
type MemoryRepository struct {
	orders map[string]entity.Order
	mx     *sync.RWMutex
}

// New is a factory function to generate a new repository of orders
func New() *MemoryRepository {
	return &MemoryRepository{
		orders: make(map[string]entity.Order),
		mx:     &sync.RWMutex{},
	}
}

// Get finds a order by ID
func (mr *MemoryRepository) Get(id string) (entity.Order, bool) {
	mr.mx.RLock()
	o, ok := mr.orders[id]
	mr.mx.RUnlock()

	return o, ok
}

// Add will add a new order to the repository
func (mr *MemoryRepository) Add(o *entity.Order) error {
	mr.mx.Lock()
	defer mr.mx.Unlock()

	if mr.orders == nil {
		// Safety check if order is not create
		mr.orders = make(map[string]entity.Order)
	}
	// Make sure order isn't already in the repository
	o.ID = uuid.NewString()
	if _, ok := mr.orders[o.ID]; ok {
		return fmt.Errorf("order already exists: %w", order.ErrFailedToAddOrder)
	}
	mr.orders[o.ID] = *o

	return nil
}
