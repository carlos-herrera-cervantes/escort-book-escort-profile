package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/iidentification_repository.go -package=mocks --build_flags=--mod=mod . IIdentificationRepository
type IIdentificationRepository interface {
	GetAll(ctx context.Context, profileId string) ([]models.Identification, error)
	GetOne(ctx context.Context, id string) (models.Identification, error)
	Create(ctx context.Context, identification *models.Identification) error
	UpdateOne(ctx context.Context, id string, identification *models.Identification) error
	Count(ctx context.Context) (int, error)
}
