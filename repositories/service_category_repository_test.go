package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestServiceCategoryRepositoryGetAll(t *testing.T) {
	serviceCategoryRepository := ServiceCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := serviceCategoryRepository.GetAll(ctxWithCancel, 1, 10)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		categories, err := serviceCategoryRepository.GetAll(
			context.Background(),
			1, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, categories)
	})
}

func TestServiceCategoryRepositoryGetById(t *testing.T) {
	serviceCategoryRepository := ServiceCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := serviceCategoryRepository.GetById(
			context.Background(),
			"639ccda97b90f45d6761c1c6",
		)
		assert.Error(t, err)
	})
}

func TestServiceCategoryRepositoryCount(t *testing.T) {
	serviceCategoryRepository := ServiceCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := serviceCategoryRepository.Count(context.Background())
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
