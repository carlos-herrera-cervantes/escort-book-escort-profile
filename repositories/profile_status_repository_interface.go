package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IProfileStatusRepository interface {
	GetOne(ctx context.Context, id string) (models.ProfileStatus, error)
	Create(ctx context.Context, profileStatus *models.ProfileStatus) error
	UpdateOne(ctx context.Context, id string, profileStatus *models.ProfileStatus) error
}
