package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestServiceRepositoryGetAll(t *testing.T) {
	serviceRepository := ServiceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := serviceRepository.GetAll(
			ctxWithCancel,
			"639cd0d42fa37dd4d02843a5",
			1, 10,
		)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		services, err := serviceRepository.GetAll(
			context.Background(),
			"639cd0d42fa37dd4d02843a5",
			1, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, services)
	})
}

func TestServiceRepositoryGetOne(t *testing.T) {
	serviceRepository := ServiceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := serviceRepository.GetOne(
			context.Background(),
			"639cd21d510c0e88c9bc07b9",
			"639cd23a8a41d56ffddbe638",
		)
		assert.Error(t, err)
	})
}

func TestServiceRepositoryCreate(t *testing.T) {
	serviceRepository := ServiceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := serviceRepository.Create(context.Background(), &models.Service{})
		assert.Error(t, err)
	})
}

func TestServiceRepositoryDeleteOne(t *testing.T) {
	serviceRepository := ServiceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := serviceRepository.DeleteOne(
			ctxWithCancel,
			"639cd21d510c0e88c9bc07b9",
			"639cd23a8a41d56ffddbe638",
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := serviceRepository.DeleteOne(
			context.Background(),
			"639cd21d510c0e88c9bc07b9",
			"639cd23a8a41d56ffddbe638",
		)
		assert.NoError(t, err)
	})
}

func TestServiceRepositoryCount(t *testing.T) {
	serviceRepository := ServiceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := serviceRepository.Count(
			context.Background(),
			"639cd21d510c0e88c9bc07b9",
		)
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
