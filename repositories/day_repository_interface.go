package repositories

import (
	"context"
	"escort-book-escort-profile/models"
)

type IDayRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.Day, error)
	GetOneByName(ctx context.Context, name string) (models.Day, error)
	GetById(ctx context.Context, id string) (models.Day, error)
	Count(ctx context.Context) (int, error)
}
