package repositories

import (
	"context"

	"escort-book-escort-profile/models"
)

//go:generate mockgen -destination=./mocks/ischedule_repository.go -package=mocks --build_flags=--mod=mod . IScheduleRepository
type IScheduleRepository interface {
	GetAll(ctx context.Context, profileId string, offset, limit int) ([]models.Schedule, error)
	GetOne(ctx context.Context, id string) (models.Schedule, error)
	Create(ctx context.Context, schedule *models.Schedule) error
	DeleteOne(ctx context.Context, id string) error
	Count(ctx context.Context, profileId string) (int, error)
}
