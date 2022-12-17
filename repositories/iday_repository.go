package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/iday_repository.go -package=mocks --build_flags=--mod=mod . IDayRepository
type IDayRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.Day, error)
	GetOneByName(ctx context.Context, name string) (models.Day, error)
	GetById(ctx context.Context, id string) (models.Day, error)
	Count(ctx context.Context) (int, error)
}
