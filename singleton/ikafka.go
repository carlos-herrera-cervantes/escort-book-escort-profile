package singleton

import "github.com/confluentinc/confluent-kafka-go/kafka"

//go:generate mockgen -destination=./mocks/ikafka.go -package=mocks --build_flags=--mod=mod . IKafka
type IKafka interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
}
