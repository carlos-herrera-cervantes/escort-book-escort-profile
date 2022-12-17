package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestAttentionSiteCategoryRepositorygGetAll(t *testing.T) {
	attentionSiteCategoryRepository := AttentionSiteCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return empty slice when there are no rows", func(t *testing.T) {
		categories, err := attentionSiteCategoryRepository.GetAll(context.Background(), 1, 10)
		assert.NoError(t, err)
		assert.Empty(t, categories)
	})

	t.Run("Should return error when category does not exists", func(t *testing.T) {
		_, err := attentionSiteCategoryRepository.GetById(
			context.Background(),
			"82b34257-96a9-428c-9bba-9b98de3edba4",
		)
		assert.Error(t, err)
	})

	t.Run("Should return 0 rows", func(t *testing.T) {
		counter, err := attentionSiteCategoryRepository.Count(context.Background())
		assert.Zero(t, counter)
		assert.NoError(t, err)
	})
}
