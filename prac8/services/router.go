package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"prac8/errorHandling"
	"prac8/utilTypes"
)

func (service *OrderService) ConfigureRouter() *mux.Router {
	return Handle(service)
}

func Handle(service *OrderService) *mux.Router {
	router := mux.Router{}
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		order := utilTypes.OrderDto{}
		order.OrderId = 2
		order.SellerId = 2
		orderJson, err := json.Marshal(order)
		errorHandling.HandleError(err, "")
		service.CreateOrder(orderJson)
	}).Methods(http.MethodGet)

	router.HandleFunc("/seller", func(writer http.ResponseWriter, request *http.Request) {
		order := utilTypes.OrderDto{}
		order.OrderId = -2
		order.SellerId = 2
		orderJson, err := json.Marshal(order)
		errorHandling.HandleError(err, "")
		service.CreateOrder(orderJson)
	}).Methods(http.MethodGet)

	router.HandleFunc("/payment", func(writer http.ResponseWriter, request *http.Request) {
		order := utilTypes.OrderDto{}
		order.OrderId = 2
		order.SellerId = -2
		orderJson, err := json.Marshal(order)
		errorHandling.HandleError(err, "")
		service.CreateOrder(orderJson)
	}).Methods(http.MethodGet)
	return &router
}
