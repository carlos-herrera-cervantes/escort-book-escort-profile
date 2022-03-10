package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IPriceRepository interface {
	GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Price, error)
	GetOne(ctx context.Context, id string) (models.Price, error)
	Create(ctx context.Context, price *models.Price) error
	UpdateOne(ctx context.Context, id string, price *models.Price) error
	DeleteOne(ctx context.Context, id string) error
	Count(ctx context.Context, profileId string) (int, error)
}
