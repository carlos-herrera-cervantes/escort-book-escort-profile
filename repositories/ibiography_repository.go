package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/ibiography_repository.go -package=mocks --build_flags=--mod=mod . IBiographyRepository
type IBiographyRepository interface {
	GetOne(ctx context.Context, id string) (models.Biography, error)
	Create(ctx context.Context, biography *models.Biography) error
	UpdateOne(ctx context.Context, id string, biography *models.Biography) error
	DeleteOne(ctx context.Context, id string) error
}
