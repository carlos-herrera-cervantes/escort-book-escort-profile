package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IAvatarRepository interface {
	GetOne(ctx context.Context, id string) (models.Avatar, error)
	Create(ctx context.Context, avatar *models.Avatar) error
	UpdateOne(ctx context.Context, id string, avatar *models.Avatar) error
	DeleteOne(ctx context.Context, id string) error
}
