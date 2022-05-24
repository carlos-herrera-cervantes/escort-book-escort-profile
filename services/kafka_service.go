package services

import (
	"context"
	"encoding/json"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/types"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaService struct {
	Producer *db.Producer
}

func (k *KafkaService) SendMessage(ctx context.Context, topic string, blockUserEvent types.BlockUserEvent) error {
	value, _ := json.Marshal(blockUserEvent)

	err := k.Producer.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
