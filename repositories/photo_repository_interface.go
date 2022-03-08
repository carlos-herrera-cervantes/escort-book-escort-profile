package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IPhotoRepository interface {
	GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Photo, error)
	GetOne(ctx context.Context, id string) (models.Photo, error)
	Create(ctx context.Context, photo *models.Photo) error
	DeleteOne(ctx context.Context, id string) error
	Count(ctx context.Context, profileId string) (int, error)
}
