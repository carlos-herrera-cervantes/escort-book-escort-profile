package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IProfileRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.Profile, error)
	GetOne(ctx context.Context, id string) (models.Profile, error)
	Create(ctx context.Context, profile *models.Profile) error
	UpdateOne(ctx context.Context, id string, profile *models.Profile) error
	DeleteOne(ctx context.Context, id string) error
	Count(ctx context.Context) (int, error)
}
