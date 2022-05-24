package services

import (
	"context"
	"escort-book-escort-profile/types"
)

type IKafkaService interface {
	SendMessage(ctx context.Context, topic string, blockUserEvent types.BlockUserEvent) error
}
