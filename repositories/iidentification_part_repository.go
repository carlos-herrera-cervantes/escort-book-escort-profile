package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/iidentification_part_repository.go -package=mocks --build_flags=--mod=mod . IIdentificationPartRepository
type IIdentificationPartRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.IdentificationPart, error)
	GetById(ctx context.Context, id string) (models.IdentificationPart, error)
	Count(ctx context.Context) (int, error)
}
