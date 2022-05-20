package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IIdentificationRepository interface {
	GetAll(ctx context.Context, profileId string) ([]models.Identification, error)
	GetOne(ctx context.Context, id string) (models.Identification, error)
	Create(ctx context.Context, identification *models.Identification) error
	UpdateOne(ctx context.Context, id string, identification *models.Identification) error
	Count(ctx context.Context) (int, error)
}
