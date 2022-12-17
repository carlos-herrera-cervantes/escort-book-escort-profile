package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestTimeCategoryRepositoryGetAll(t *testing.T) {
	timeCategoryRepository := TimeCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := timeCategoryRepository.GetAll(ctxWithCancel, 1, 10)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		categories, err := timeCategoryRepository.GetAll(
			context.Background(),
			1, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, categories)
	})
}

func TestTimeCategoryRepositoryGetById(t *testing.T) {
	timeCategoryRepository := TimeCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := timeCategoryRepository.GetById(
			context.Background(),
			"639d3ed1ec1e659f9e5cddf1",
		)
		assert.Error(t, err)
	})
}

func TestTimeCategoryRepositoryCount(t *testing.T) {
	timeCategoryRepository := TimeCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := timeCategoryRepository.Count(context.Background())
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
