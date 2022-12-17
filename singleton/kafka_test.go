package singleton

import (
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestNewProducer(t *testing.T) {
	t.Run("Should return a pointer to kafka producer", func(t *testing.T) {
		producer := NewProducer()
		assert.IsType(t, &kafka.Producer{}, producer)
	})
}
