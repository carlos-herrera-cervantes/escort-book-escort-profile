package services

import (
	"context"
	"escort-book-escort-profile/db"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaService struct {
	Producer *db.Producer
}

func (k *KafkaService) SendMessage(ctx context.Context, topic string, message []byte) error {
	err := k.Producer.KafkaProducer.Produce(&kafka.Message{
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
