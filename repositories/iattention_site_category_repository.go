package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/iattention_site_category_repository.go -package=mocks --build_flags=--mod=mod . IAttentionSiteCategoryRepository
type IAttentionSiteCategoryRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.AttentionSiteCategory, error)
	GetById(ctx context.Context, id string) (models.AttentionSiteCategory, error)
	Count(ctx context.Context) (int, error)
}
