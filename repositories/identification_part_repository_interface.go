package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IIdentificationPartRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.IdentificationPart, error)
	Count(ctx context.Context) (int, error)
}
