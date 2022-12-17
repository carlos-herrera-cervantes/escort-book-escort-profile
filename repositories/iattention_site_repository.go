package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/iattention_site_repository.go -package=mocks --build_flags=--mod=mod . IAttentionSiteRepository
type IAttentionSiteRepository interface {
	GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.AttentionSiteDetailed, error)
	GetOne(ctx context.Context, id string) (models.AttentionSiteDetailed, error)
	Create(ctx context.Context, site *models.AttentionSite) error
	DeleteOne(ctx context.Context, id string) error
	Count(ctx context.Context, profileId string) (int, error)
}
