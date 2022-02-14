package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IAttentionSiteRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.AttentionSite, error)
	GetOne(ctx context.Context, id string) (models.AttentionSite, error)
	Create(ctx context.Context, site *models.AttentionSite) error
	DeleteOne(ctx context.Context, id string) error
	Count(ctx context.Context) (int, error)
}
