package services

type ServiceConsumer interface {
	StartConsuming(handler func([]byte))
	SetService(service Service)
}
