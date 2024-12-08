package service

import (
	"applicationDesignTest/domain"
	"applicationDesignTest/entity"
	"applicationDesignTest/pkg"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const Hotel = "reddison"
const Room = "lux"

type appTest struct {
	srv   *Order
	srvAv *Availability
}

func TestOrderService(t *testing.T) {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(logHandler)
	availService := NewAvailabilityService(logger)
	orderService := NewOrderService(logger, availService)

	app := new(appTest)
	app.srv = orderService
	app.srvAv = availService
	// add rooms
	err := app.AddAvail(Hotel, Room, pkg.Date(2024, 1, 1), pkg.Date(2024, 2, 1))
	if assert.NoError(t, err, "adding rooms 1") {
		list0 := app.srvAv.GetRoomAvailabilityList(Hotel, Room)
		assert.Len(t, list0, 32)
	}

	// add rooms
	err = app.AddAvail(Hotel, Room, pkg.Date(2024, 1, 1), pkg.Date(2024, 2, 1))
	if assert.NoError(t, err, "adding rooms 2") {
		list1 := app.srvAv.GetRoomAvailabilityList(Hotel, Room)
		assert.Len(t, list1, 32)
	}

	order1 := &entity.Order{HotelID: Hotel, RoomID: Room, UserEmail: "test1@mail.ru"}
	order1.From = pkg.Date(2024, 1, 1)
	order1.To = pkg.Date(2024, 1, 31)
	err = app.srv.CreateOrder(order1)
	if assert.NoError(t, err, "add order") {
		o1, ok1 := app.srv.GetOrder(order1.ID)
		assert.True(t, ok1)
		assert.Equal(t, o1, *order1, "order1 equal")
	}

	list2 := app.srvAv.GetRoomAvailabilityList(Hotel, Room)
	assert.Len(t, list2, 32)

	order2 := &entity.Order{HotelID: Hotel, RoomID: Room, UserEmail: "test2@mail.ru"}
	order2.From = pkg.Date(2024, 1, 31)
	order2.To = pkg.Date(2024, 1, 1)
	err = app.srv.CreateOrder(order2)
	if assert.NoError(t, err, "add order 2") {
		o2, ok2 := app.srv.GetOrder(order2.ID)
		assert.True(t, ok2)
		assert.Equal(t, o2, *order2, "order2 equal")
	}

	list3 := app.srvAv.GetRoomAvailabilityList(Hotel, Room)
	assert.Len(t, list3, 1)

	order3 := &entity.Order{HotelID: Hotel, RoomID: Room, UserEmail: "test3@mail.ru"}
	order3.From = pkg.Date(2024, 2, 1)
	order3.To = pkg.Date(2024, 2, 14)

	err = app.srv.CreateOrder(order3)
	if assert.ErrorIs(t, err, domain.ErrIntervalNotAvailableForBooking) {
		_, ok3 := app.srv.GetOrder(order3.ID)
		assert.False(t, ok3)
	}

	order4 := &entity.Order{HotelID: Hotel, RoomID: Room, UserEmail: "test4@mail.ru"}
	order4.From = pkg.Date(2024, 2, 1)
	order4.To = pkg.Date(2024, 1, 14)
	//todo: is error?
	err = app.srv.CreateOrder(order4)
	app.srv.log.Error("??? is error", "err", err)

	_, ok4 := app.srv.GetOrder(order4.ID)

	assert.False(t, ok4)
}

func (app *appTest) AddAvail(hotel, room string, from, to time.Time) error {
	days := &entity.DateInterval{From: from, To: to}
	i := days.GetInterval()
	for _, d := range i {
		if err := app.srvAv.AddRoomAvailability(entity.RoomAvailability{HotelID: hotel, RoomID: room, Date: d, Quota: 1}); err != nil {
			return err
		}
	}

	return nil
}
