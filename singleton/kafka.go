package singleton

import (
	"log"
	"sync"

	"escort-book-escort-profile/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer
var singleKafkaClient sync.Once

func initProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.InitKafka().BootstrapServers,
		"client.id":         config.InitKafka().GroupId,
		"acks":              "all",
	})

	if err != nil {
		log.Panic("ERROR CREATING A PRODUCER: ", err)
	}

	producer = p
}

func NewProducer() *kafka.Producer {
	singleKafkaClient.Do(initProducer)
	return producer
}
