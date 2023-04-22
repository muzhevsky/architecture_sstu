package services

type Service interface {
	OnConsume(message []byte)
	Start()
}
