package services

import (
	"encoding/json"
	"fmt"
	"prac8/errorHandling"
	"prac8/producer"
	"prac8/utilTypes"
)

type PaymentService struct {
	consumer *ServiceConsumer
	producer *producer.ServiceProducer
}

func (s *PaymentService) AcceptOrder(order []byte) {
	(*(s.producer)).SendMessage(order)
}

func NewPayment(producer producer.ServiceProducer, consumer ServiceConsumer) (service *PaymentService) {
	return &PaymentService{
		consumer: &consumer,
		producer: &producer,
	}
}

func (s *PaymentService) Reply(order []byte) {
	(*(s.producer)).SendMessage(order)
}

func (s *PaymentService) Start() {
	go (*(s.consumer)).StartConsuming(func(a []byte) {})
}

func (s *PaymentService) OnConsume(message []byte) {
	mes := &utilTypes.OrderDto{}
	err := json.Unmarshal(message, mes)
	errorHandling.HandleError(err, "")

	if mes.SellerId < 0 {
		mes.Status = int(utilTypes.PaymentRejected)
		fmt.Println("payment rejected")
	} else {
		mes.Status = int(utilTypes.PaymentSucceeded)
		fmt.Println("payment succeeded")
	}

	result, err := json.Marshal(mes)
	errorHandling.HandleError(err, "")

	(*(s.producer)).SendMessage(result)
}
