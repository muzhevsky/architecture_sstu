package services

import (
	"encoding/json"
	"fmt"
	"prac8/errorHandling"
	"prac8/producer"
	"prac8/utilTypes"
)

type OrderService struct {
	sellerProducer  producer.ServiceProducer
	paymentProducer producer.ServiceProducer

	answerConsumer ServiceConsumer
}

func NewOrderService(connectionString string) *OrderService {
	service := &OrderService{}

	service.sellerProducer = producer.NewDefault(connectionString, "seller")
	service.paymentProducer = producer.NewDefault(connectionString, "payment")

	service.answerConsumer = NewDefault(connectionString, "answer")

	return service
}

func (service *OrderService) CreateOrder(message []byte) {
	fmt.Println("\nTrying to create order")
	service.sellerProducer.SendMessage(message)
}

func (service *OrderService) Start() {
	service.answerConsumer.SetService(service)
	go service.answerConsumer.StartConsuming(func(a []byte) {})
}

func (service *OrderService) OnConsume(message []byte) {
	mes := &utilTypes.OrderDto{}
	err := json.Unmarshal(message, mes)
	errorHandling.HandleError(err, "")

	switch mes.Status {
	case int(utilTypes.SellerRejected):
		fmt.Println("order cancelled")
		break
	case int(utilTypes.PaymentRejected):
		mes.Status = int(utilTypes.SellerRejected)
		reply, err := json.Marshal(mes)
		errorHandling.HandleError(err, "")
		service.sellerProducer.SendMessage(reply)
		break
	case int(utilTypes.PaymentSucceeded):
		fmt.Println("order created")
		break
	case int(utilTypes.SellerSucceeded):
		reply, err := json.Marshal(mes)
		errorHandling.HandleError(err, "")
		service.paymentProducer.SendMessage(reply)
	}
}
