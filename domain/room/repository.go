package room

import (
	"applicationDesignTest/entity"
	"errors"
	"time"
)

var (
	// ErrFailedToAddRoomAvailability is returned when RoomAvailability could not be added to the repository.
	ErrFailedToAddRoomAvailability = errors.New("failed to add RoomAvailability to the repository")
)

type RoomAvailabilityRepository interface {
	Add(...entity.RoomAvailability) error
	Get(hotel, room string, date time.Time) (ra entity.RoomAvailability, ok bool)
	GetList(hotel, room string) []entity.RoomAvailability
}
