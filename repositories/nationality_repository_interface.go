package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type INationalityRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.Nationality, error)
	GetById(ctx context.Context, id string) (models.Nationality, error)
	Count(ctx context.Context) (int, error)
}
