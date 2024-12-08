package service

import (
	"applicationDesignTest/config"
	"applicationDesignTest/domain"
	"applicationDesignTest/domain/order"
	orderMem "applicationDesignTest/domain/order/memory"
	"applicationDesignTest/entity"
)

type OrderService interface {
	CreateOrder(*entity.Order) error
	GetOrder(id string) (entity.Order, bool)
}

// OrderService is a implementation of the OrderService
type Order struct {
	orders       order.OrderRepository
	availability AvailabilityService
	log          config.Logger
}

func NewOrderService() *Order {
	// Create the order service
	os := new(Order)

	return os
}

// WithMemoryOrderRepository adds a in memory product repo
func (o *Order) WithMemoryOrderRepository() *Order {
	// Create the memory repo
	o.orders = orderMem.New()

	return o
}

// WithMemoryRoomAvailabilityRepository adds a in memory product repo
func (o *Order) WithRoomAvailabilityService(a AvailabilityService) *Order {
	// Create the memory repo
	o.availability = a

	return o
}

// WithLogger adds logger to service
func (o *Order) WithLogger(l config.Logger) *Order {
	// Create the memory repo
	o.log = l

	return o
}

// CreateOrder will chaintogether all repositories to create order for customer
func (o *Order) CreateOrder(no *entity.Order) error {
	if no.From.After(no.To) {
		no.From, no.To = no.To, no.From
	}

	orderDays := no.GetInterval()
	o.log.Debug("order create", "o", no)
	if len(orderDays) == 0 {
		return domain.ErrEmptyBookingInterval
	}

	available := make([]entity.RoomAvailability, 0)
	for _, od := range orderDays {
		if ra, ok := o.availability.GetRoomAvailability(no.HotelID, no.RoomID, od); ok {
			ra.Quota = -1
			available = append(available, ra)
		} else {
			return domain.ErrIntervalNotAvailableForBooking
		}
	}

	if err := o.availability.AddRoomAvailability(available...); err != nil {
		return err
	}

	err := o.orders.Add(no)
	if err != nil {
		if err := o.availability.AddRoomAvailability(available...); err != nil {
			return err
		}
	}

	o.log.Debug("Order successfully created", "order", *no)

	return nil
}

// GetOrder will get order from repository
func (o *Order) GetOrder(id string) (entity.Order, bool) {
	order, ok := o.orders.Get(id)
	return order, ok
}
