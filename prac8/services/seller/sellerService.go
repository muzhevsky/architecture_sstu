package seller

import (
	"prac8/consumer"
	"prac8/producer"
)

type sellerService struct {
	consumer *consumer.ServiceConsumer
	producer *producer.ServiceProducer
}

func (s *sellerService) AcceptOrder(order []byte) {
	(*(s.producer)).SendMessage(order)
}

func New(producer producer.ServiceProducer, consumer consumer.ServiceConsumer) (service *sellerService) {
	return &sellerService{
		consumer: &consumer,
		producer: &producer,
	}
}
