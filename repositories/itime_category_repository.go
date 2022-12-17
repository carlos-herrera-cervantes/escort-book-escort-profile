package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/itime_category_repository.go -package=mocks --build_flags=--mod=mod . ITimeCategoryRepository
type ITimeCategoryRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.TimeCategory, error)
	GetById(ctx context.Context, id string) (models.TimeCategory, error)
	Count(ctx context.Context) (int, error)
}
