package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestAttentionSiteRepositoryGetAll(t *testing.T) {
	attentionSiteRepository := AttentionSiteRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return empty slice", func(t *testing.T) {
		sites, err := attentionSiteRepository.GetAll(
			context.Background(),
			"6398bf735659201469997bc0",
			1, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, sites)
	})
}

func TestAttentionSiteRepositoryGetOne(t *testing.T) {
	attentionSiteRepository := AttentionSiteRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error when row does not exists", func(t *testing.T) {
		_, err := attentionSiteRepository.GetOne(
			context.Background(),
			"94247409-40db-4c13-878f-9d21861c5506",
		)
		assert.Error(t, err)
	})
}

func TestAttentionSiteRepositoryDeleteOne(t *testing.T) {
	attentionSiteRepository := AttentionSiteRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := attentionSiteRepository.DeleteOne(
			ctxWithCancel,
			"6398d43b7f3427b956ec4b44",
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := attentionSiteRepository.DeleteOne(
			context.Background(),
			"6398d43b7f3427b956ec4b44",
		)
		assert.NoError(t, err)
	})
}

func TestAttentionSiteRepositoryCount(t *testing.T) {
	attentionSiteRepository := AttentionSiteRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := attentionSiteRepository.Count(
			context.Background(),
			"6398d43b7f3427b956ec4b44",
		)
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
