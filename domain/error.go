package domain

import "errors"

var ErrEmptyBookingInterval = errors.New("empty booking interval")
var ErrIntervalNotAvailableForBooking = errors.New("interval is not available for booking")
