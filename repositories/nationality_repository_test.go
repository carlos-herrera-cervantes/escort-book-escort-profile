package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestNationalityRepositoryGetAll(t *testing.T) {
	nationalityRepository := NationalityRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := nationalityRepository.GetAll(ctxWithCancel, 1, 10)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		nationalities, err := nationalityRepository.GetAll(context.Background(), 1, 10)
		assert.NoError(t, err)
		assert.Empty(t, nationalities)
	})
}

func TestNationalityRepositoryGetById(t *testing.T) {
	nationalityRepository := NationalityRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := nationalityRepository.GetById(
			context.Background(),
			"639c04c84b5e054772a16df8",
		)
		assert.Error(t, err)
	})
}

func TestNationalityRepositoryCount(t *testing.T) {
	nationalityRepository := NationalityRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := nationalityRepository.Count(context.Background())
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
