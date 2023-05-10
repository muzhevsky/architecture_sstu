package services

import (
	"encoding/json"
	"fmt"
	"prac8/errorHandling"
	"prac8/producer"
	"prac8/utilTypes"
)

type SellerService struct {
	consumer ServiceConsumer
	producer producer.ServiceProducer
}

func (s *SellerService) AcceptOrder(order []byte) {
	s.producer.SendMessage(order)
}

func NewSeller(producer producer.ServiceProducer, consumer ServiceConsumer) (service *SellerService) {
	return &SellerService{
		consumer: consumer,
		producer: producer,
	}
}

func (s *SellerService) Reply(order []byte) {
	s.producer.SendMessage(order)
}

func (s *SellerService) Start() {
	go s.consumer.StartConsuming(func(a []byte) {})
}

func (s *SellerService) OnConsume(message []byte) {
	mes := &utilTypes.OrderDto{}
	err := json.Unmarshal(message, mes)
	errorHandling.HandleError(err, "")

	if mes.Status == int(utilTypes.SellerRejected) {
		fmt.Println("seller cancelled")
		result, err := json.Marshal(mes)
		errorHandling.HandleError(err, "")
		s.producer.SendMessage(result)
		return
	}

	if mes.OrderId < 0 {
		mes.Status = int(utilTypes.SellerRejected)
		fmt.Println("seller rejected")
	} else {
		mes.Status = int(utilTypes.SellerSucceeded)
		fmt.Println("seller succeeded")
	}

	result, err := json.Marshal(mes)
	errorHandling.HandleError(err, "")

	s.producer.SendMessage(result)
}
