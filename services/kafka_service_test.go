package services

import (
	"errors"
	"testing"

	mockSingleton "escort-book-escort-profile/singleton/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestKafkaServiceSendMessage(t *testing.T) {
	controller := gomock.NewController(t)
	mockKafka := mockSingleton.NewMockIKafka(controller)
	kafkaService := KafkaService{Producer: mockKafka}

	t.Run("Should return error when Produce fails", func(t *testing.T) {
		mockKafka.
			EXPECT().
			Produce(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		topic := "test-topic"
		message := []byte("hello world")
		err := kafkaService.SendMessage(topic, message)

		assert.Error(t, err)
	})

	t.Run("Should return nil when Produce succeeds", func(t *testing.T) {
		mockKafka.
			EXPECT().
			Produce(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		topic := "test-topic"
		message := []byte("hello world")
		err := kafkaService.SendMessage(topic, message)

		assert.NoError(t, err)
	})
}
