package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/iservice_category_repository.go -package=mocks --build_flags=--mod=mod . IServiceCategoryRepository
type IServiceCategoryRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.ServiceCategory, error)
	GetById(ctx context.Context, id string) (models.ServiceCategory, error)
	Count(ctx context.Context) (int, error)
}
