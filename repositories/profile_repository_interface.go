package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IProfileRepository interface {
	GetOne(ctx context.Context, id string) (models.Profile, error)
}
