package order

import (
	"applicationDesignTest/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type OrderGet struct {
	srv service.OrderService
}

func NewOrderGet(srv service.OrderService) *OrderGet {
	handler := new(OrderGet)
	handler.srv = srv
	return handler
}

func (s *OrderGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "ID")
	if orderID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, ok := s.srv.GetOrder(orderID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if resp, err := json.Marshal(order); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}
