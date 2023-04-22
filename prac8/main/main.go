package main

import (
	"net/http"
	"prac8/producer"
	"prac8/services"
)

type person struct {
	Name string
	Age  int
}

func main() {
	initServices("amqp://guest:guest@localhost:5672/")
}

func initServices(connectionString string) {
	initSeller(connectionString)
	initPayment(connectionString)

	orderService := services.NewOrderService(connectionString)
	orderService.Start()
	router := orderService.ConfigureRouter()
	http.ListenAndServe("127.0.0.1:8080", router)
}

func initPayment(connectionString string) {
	producer := producer.NewDefault(connectionString, "answer")
	consumer := services.NewDefault(connectionString, "payment")

	paymentService := services.NewPayment(producer, consumer)
	consumer.SetService(paymentService)
	paymentService.Start()
}

func initSeller(connectionString string) {
	producer := producer.NewDefault(connectionString, "answer")
	consumer := services.NewDefault(connectionString, "seller")

	sellerService := services.NewSeller(producer, consumer)
	consumer.SetService(sellerService)
	sellerService.Start()
}
