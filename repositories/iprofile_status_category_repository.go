package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/iprofile_status_category_repository.go -package=mocks --build_flags=--mod=mod . IProfileStatusCategoryRepository
type IProfileStatusCategoryRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.ProfileStatusCategory, error)
	GetOneByName(ctx context.Context, name string) (models.ProfileStatusCategory, error)
	GetById(ctx context.Context, id string) (models.ProfileStatusCategory, error)
	Count(ctx context.Context) (int, error)
}
