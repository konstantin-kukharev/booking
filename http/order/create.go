package order

import (
	"applicationDesignTest/entity"
	"applicationDesignTest/service"
	"encoding/json"
	"io"
	"net/http"
)

type OrderCreate struct {
	srv service.OrderService
}

func NewOrderCreate(srv service.OrderService) *OrderCreate {
	handler := new(OrderCreate)
	handler.srv = srv

	return handler
}

func (s *OrderCreate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	order := new(entity.Order)
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = order.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.srv.CreateOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}
