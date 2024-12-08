package service

import (
	"applicationDesignTest/config"
	"applicationDesignTest/domain/room"
	roomMem "applicationDesignTest/domain/room/memory"
	"applicationDesignTest/entity"
	"time"
)

type AvailabilityService interface {
	AddRoomAvailability(...entity.RoomAvailability) error
	GetRoomAvailability(hotel, room string, date time.Time) (entity.RoomAvailability, bool)
	GetRoomAvailabilityList(hotel, room string) []entity.RoomAvailability
}

// Availability is a implementation of the AvailabilityService
type Availability struct {
	availability room.RoomAvailabilityRepository
	log          config.Logger
}

func NewAvailabilityService(l config.Logger) *Availability {
	// Create the availability service
	as := new(Availability).
		WithMemoryRoomAvailabilityRepository().
		WithLogger(l)

	return as
}

// WithMemoryRoomAvailabilityRepository adds a in memory room availability repo
func (a *Availability) WithMemoryRoomAvailabilityRepository() *Availability {
	// Create the memory repo
	a.availability = roomMem.New()

	return a
}

// WithLogger adds logger to service
func (a *Availability) WithLogger(l config.Logger) *Availability {
	a.log = l

	return a
}

// AddRoomAvailability add room availability
func (a *Availability) AddRoomAvailability(ra ...entity.RoomAvailability) error {
	// Add orders to repo
	a.log.Debug("add room availability", "intervals", ra)
	for _, avail := range ra {
		if err := a.availability.Add(avail); err != nil {
			return err
		}
	}

	return nil
}

// GetRoomAvailability get availability by hotel id & room id & date
func (a *Availability) GetRoomAvailability(h, r string, d time.Time) (entity.RoomAvailability, bool) {
	// Add orders to repo
	a.log.Debug("get room availability", "hotel", h, "room", r, "date", d)
	return a.availability.Get(h, r, d)
}

// GetRoomAvailabilityList get availability list by hotel id & room id
func (a *Availability) GetRoomAvailabilityList(h, r string) []entity.RoomAvailability {
	// Add orders to repo
	a.log.Debug("get room availability", "hotel", h, "room", r)
	return a.availability.GetList(h, r)
}
