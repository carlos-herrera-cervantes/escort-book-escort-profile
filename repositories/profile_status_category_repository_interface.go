package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IProfileStatusCategoryRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.ProfileStatusCategory, error)
	GetOneByName(ctx context.Context, name string) (models.ProfileStatusCategory, error)
	Count(ctx context.Context) (int, error)
}
