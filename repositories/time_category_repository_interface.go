package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type ITimeCategoryRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.TimeCategory, error)
	Count(ctx context.Context) (int, error)
}
