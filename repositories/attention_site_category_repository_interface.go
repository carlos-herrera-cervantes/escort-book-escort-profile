package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IAttentionSiteCategoryRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.AttentionSiteCategory, error)
	GetById(ctx context.Context, id string) (models.AttentionSiteCategory, error)
	Count(ctx context.Context) (int, error)
}
