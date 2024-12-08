package main

import (
	"applicationDesignTest/config"
	availabilityHandler "applicationDesignTest/http/availability"
	orderHandler "applicationDesignTest/http/order"
	"applicationDesignTest/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	app := config.New().WithDebug()
	err := run(app)

	app.Log().Error("server exited with error", "err", err)
}

func run(app config.Application) error {
	availService := service.NewAvailabilityService(app.Log())
	orderService := service.NewOrderService(app.Log(), availService)

	hOrderAdd := orderHandler.NewOrderCreate(orderService)
	hOrderGet := orderHandler.NewOrderGet(orderService)
	hAvailAdd := availabilityHandler.NewAvailabilityAdd(availService)
	hAvailGet := availabilityHandler.NewAvailabilityGet(availService)

	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		r.Post("/orders", hOrderAdd.ServeHTTP)
		r.Post("/availability", hAvailAdd.ServeHTTP)
		r.Get("/orders/{ID}", hOrderGet.ServeHTTP)
		r.Get("/availability/{HotelID}/{RoomID}", hAvailGet.ServeHTTP)
	})

	app.Log().Info("Server listening on", "address", app.GetServerAddress())

	return http.ListenAndServe(app.GetServerAddress(), r)
}
