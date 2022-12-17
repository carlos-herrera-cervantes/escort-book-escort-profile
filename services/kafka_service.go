package services

import (
	"escort-book-escort-profile/singleton"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaService struct {
	Producer singleton.IKafka
}

func (k *KafkaService) SendMessage(topic string, message []byte) error {
	err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: message,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
