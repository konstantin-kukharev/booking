package order

import (
	"applicationDesignTest/entity"
	"applicationDesignTest/service"
	"encoding/json"
	"io"
	"net/http"
)

type AvailabilityAdd struct {
	srv service.AvailabilityService
}

func NewAvailabilityAdd(srv service.AvailabilityService) *AvailabilityAdd {
	handler := new(AvailabilityAdd)
	handler.srv = srv

	return handler
}

func (s *AvailabilityAdd) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	availability := new(entity.RoomAvailability)
	req, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(req, availability)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.srv.AddRoomAvailability(*availability)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(req)
}
