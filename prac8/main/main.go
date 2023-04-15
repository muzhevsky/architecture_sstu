package main

import (
	"prac8/consumer"
	"prac8/producer"
	"prac8/services/seller"
)

func main() {
	producer := producer.NewDefault("amqp://guest:guest@localhost:5672/", "testIn")
	consumer := consumer.NewDefault("amqp://guest:guest@localhost:5672/", "testOut")

	go consumer.StartConsuming()

	sellerService := seller.New(producer, consumer)

	slice := make([]byte, 100)
	for i := 0; i < 100; i++ {
		slice[i] = byte(i) << 1
	}
	sellerService.AcceptOrder(slice)
}
