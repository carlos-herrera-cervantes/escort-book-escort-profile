package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestIdentificationRepositoryGetAll(t *testing.T) {
	identificationRepository := IdentificationRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return empty slice", func(t *testing.T) {
		identifications, err := identificationRepository.GetAll(
			context.Background(),
			"639ac12e4c6c4a2a37e83ffd",
		)
		assert.Empty(t, identifications)
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := identificationRepository.GetAll(
			ctxWithCancel,
			"639ac12e4c6c4a2a37e83ffd",
		)
		assert.Error(t, err)
	})
}

func TestIdentificationRepositoryGetOne(t *testing.T) {
	identificationRepository := IdentificationRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := identificationRepository.GetOne(
			context.Background(),
			"639ac28777478218d88395ad",
		)
		assert.Error(t, err)
	})
}

func TestIdentificationRepositoryCreate(t *testing.T) {
	identificationRepository := IdentificationRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := identificationRepository.Create(
			context.Background(),
			&models.Identification{},
		)
		assert.Error(t, err)
	})
}

func TestIdentificationRepositoryUpdateOne(t *testing.T) {
	identificationRepository := IdentificationRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := identificationRepository.UpdateOne(
			ctxWithCancel,
			"639ac390231cd246e3ec415f",
			&models.Identification{},
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := identificationRepository.UpdateOne(
			context.Background(),
			"639ac390231cd246e3ec415f",
			&models.Identification{},
		)
		assert.NoError(t, err)
	})
}

func TestIdentificationRepositoryCount(t *testing.T) {
	identificationRepository := IdentificationRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := identificationRepository.Count(context.Background())
		assert.Zero(t, counter)
		assert.NoError(t, err)
	})
}
