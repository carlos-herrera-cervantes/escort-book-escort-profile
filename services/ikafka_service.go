package services

//go:generate mockgen -destination=./mocks/ikafka_service.go -package=mocks --build_flags=--mod=mod . IKafkaService
type IKafkaService interface {
	SendMessage(topic string, message []byte) error
}
