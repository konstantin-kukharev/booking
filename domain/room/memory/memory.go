package memory

import (
	"applicationDesignTest/domain/room"
	"applicationDesignTest/entity"
	"fmt"
	"sync"
	"time"
)

type key struct {
	hotel, room string
}

type interval struct {
	Data map[time.Time]int
}

func newInterval() interval {
	return interval{
		Data: map[time.Time]int{},
	}
}

// MemoryRepository fulfills the RoomAvailabilityRepository interface
type MemoryRepository struct {
	rooms map[key]interval
	sync.RWMutex
}

// New is a factory function to generate a new repository of orders
func New() *MemoryRepository {
	return &MemoryRepository{
		rooms: make(map[key]interval),
	}
}

// Get finds RoomAvailability by hotel, room, date
func (mr *MemoryRepository) Get(h, r string, d time.Time) (entity.RoomAvailability, bool) {
	mr.RLock()
	defer mr.RUnlock()
	i, ok := mr.rooms[key{hotel: h, room: r}]
	if !ok {
		return entity.RoomAvailability{}, false
	}

	if _, nq := i.Data[d]; nq {
		return entity.RoomAvailability{HotelID: h, RoomID: r, Date: d, Quota: 1}, ok
	}

	return entity.RoomAvailability{}, false
}

// GetList finds list of RoomAvailability by hotel and room
func (mr *MemoryRepository) GetList(h, r string) []entity.RoomAvailability {
	list := make([]entity.RoomAvailability, 0)
	mr.RLock()
	defer mr.RUnlock()
	if i, ok := mr.rooms[key{hotel: h, room: r}]; ok {
		for t, q := range i.Data {
			list = append(list, entity.RoomAvailability{HotelID: h, RoomID: r, Date: t, Quota: q})
		}
	}

	return list
}

// Add will add a new room(s) to the repository
func (mr *MemoryRepository) Add(av ...entity.RoomAvailability) error {
	if mr.rooms == nil {
		// Saftey check if room is not create
		mr.rooms = make(map[key]interval)
	}

	for _, ra := range av {
		k := key{hotel: ra.HotelID, room: ra.RoomID}
		currentDay := time.Now()
		if currentDay.Before(ra.Date) {
			return fmt.Errorf("room availability already exists: %w", room.ErrFailedToAddRoomAvailability)
		}

		if _, ok := mr.rooms[k]; !ok {
			mr.rooms[k] = newInterval()
		}

		mr.rooms[k].Data[ra.Date] += ra.Quota
		if mr.rooms[k].Data[ra.Date] <= 0 {
			delete(mr.rooms[k].Data, ra.Date)
		}
	}

	return nil
}
