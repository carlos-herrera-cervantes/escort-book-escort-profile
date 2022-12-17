package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestScheduleRepositoryGetAll(t *testing.T) {
	scheduleRepository := ScheduleRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := scheduleRepository.GetAll(
			ctxWithCancel,
			"639cc032f5f77eacd65d1a2f",
			1, 10,
		)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		schedules, err := scheduleRepository.GetAll(
			context.Background(),
			"639cc032f5f77eacd65d1a2f",
			1, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, schedules)
	})
}

func TestScheduleRepositoryGetOne(t *testing.T) {
	scheduleRepository := ScheduleRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := scheduleRepository.GetOne(
			context.Background(),
			"639cc032f5f77eacd65d1a2f",
		)
		assert.Error(t, err)
	})
}

func TestScheduleRepositoryCreate(t *testing.T) {
	scheduleRepository := ScheduleRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := scheduleRepository.Create(context.Background(), &models.Schedule{})
		assert.Error(t, err)
	})
}

func TestScheduleRepositoryDeleteOne(t *testing.T) {
	scheduleRepository := ScheduleRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := scheduleRepository.DeleteOne(ctxWithCancel, "639cc032f5f77eacd65d1a2f")
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := scheduleRepository.DeleteOne(
			context.Background(),
			"639cc032f5f77eacd65d1a2f",
		)
		assert.NoError(t, err)
	})
}

func TestScheduleRepositoryCount(t *testing.T) {
	scheduleRepository := ScheduleRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := scheduleRepository.Count(
			context.Background(),
			"639cc032f5f77eacd65d1a2f",
		)
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
