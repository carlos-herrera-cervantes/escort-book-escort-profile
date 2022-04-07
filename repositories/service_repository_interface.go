package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IServiceRepository interface {
	GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Service, error)
	GetOne(ctx context.Context, id, profileId string) (models.Service, error)
	Create(ctx context.Context, service *models.Service) error
	DeleteOne(ctx context.Context, id, profileId string) error
	Count(ctx context.Context, profileId string) (int, error)
}
