package producer

type ServiceProducer interface {
	SendMessage(content []byte)
}
