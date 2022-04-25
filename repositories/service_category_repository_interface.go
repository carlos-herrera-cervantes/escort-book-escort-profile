package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IServiceCategoryRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.ServiceCategory, error)
	GetById(ctx context.Context, id string) (models.ServiceCategory, error)
	Count(ctx context.Context) (int, error)
}
