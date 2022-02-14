package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IProfileRepository interface {
	GetOne(ctx context.Context, id string) (models.Profile, error)
	Create(ctx context.Context, profile *models.Profile) error
	UpdateOne(ctx context.Context, id string, profile *models.Profile) error
	DeleteOne(ctx context.Context, id string) error
}
