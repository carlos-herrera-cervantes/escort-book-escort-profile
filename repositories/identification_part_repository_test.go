package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestIdentificationPartRepositoryGetAll(t *testing.T) {
	identificationPartRepository := IdentificationPartRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return empty slice", func(t *testing.T) {
		parts, err := identificationPartRepository.GetAll(context.Background(), 1, 10)
		assert.Empty(t, parts)
		assert.NoError(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := identificationPartRepository.GetAll(ctxWithCancel, 1, 10)
		assert.Error(t, err)
	})
}

func TestIdentificationPartRepositoryGetById(t *testing.T) {
	identificationPartRepository := IdentificationPartRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := identificationPartRepository.GetById(
			context.Background(),
			"639abf462240e7c32ed9b419",
		)
		assert.Error(t, err)
	})
}

func TestIdentificationPartRepositoryCount(t *testing.T) {
	identificationPartRepository := IdentificationPartRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := identificationPartRepository.Count(context.Background())
		assert.Zero(t, counter)
		assert.NoError(t, err)
	})
}
