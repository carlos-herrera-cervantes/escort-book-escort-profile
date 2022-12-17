package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/iprofile_status_repository.go -package=mocks --build_flags=--mod=mod . IProfileStatusRepository
type IProfileStatusRepository interface {
	GetOne(ctx context.Context, id string) (models.ProfileStatus, error)
	Create(ctx context.Context, profileStatus *models.ProfileStatus) error
	UpdateOne(ctx context.Context, id string, profileStatus *models.ProfileStatus) error
}
