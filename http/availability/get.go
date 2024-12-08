package order

import (
	"applicationDesignTest/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AvailabilityGet struct {
	srv service.AvailabilityService
}

func NewAvailabilityGet(srv service.AvailabilityService) *AvailabilityGet {
	handler := new(AvailabilityGet)
	handler.srv = srv

	return handler
}

func (s *AvailabilityGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	HotelID := chi.URLParam(r, "HotelID")
	RoomID := chi.URLParam(r, "RoomID")
	if HotelID == "" || RoomID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list := s.srv.GetRoomAvailabilityList(HotelID, RoomID)

	res, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
