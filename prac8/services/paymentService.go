package services

import (
	"encoding/json"
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
	go (*(s.consumer)).StartConsuming(func(a []byte) { println("payment consumed") })
}

func (s *PaymentService) OnConsume(message []byte) {
	mes := &utilTypes.OrderDto{}
	err := json.Unmarshal(message, mes)
	errorHandling.HandleError(err, "")

	if mes.SellerId < 0 {
		mes.Status = int(utilTypes.PaymentRejected)
	} else {
		mes.Status = int(utilTypes.PaymentSucceeded)
	}

	result, err := json.Marshal(mes)
	errorHandling.HandleError(err, "")

	(*(s.producer)).SendMessage(result)
}
