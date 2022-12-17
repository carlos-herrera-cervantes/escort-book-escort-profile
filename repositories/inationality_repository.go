package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/inationality_repository.go -package=mocks --build_flags=--mod=mod . INationalityRepository
type INationalityRepository interface {
	GetAll(ctx context.Context, offset, limit int) ([]models.Nationality, error)
	GetById(ctx context.Context, id string) (models.Nationality, error)
	Count(ctx context.Context) (int, error)
}
