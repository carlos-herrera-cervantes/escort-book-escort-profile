package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestPhotoRepositoryGetAll(t *testing.T) {
	photoRepository := PhotoRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := photoRepository.GetAll(ctxWithCancel, "639c088d06b3b786e7bb9aee", 1, 10)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		photos, err := photoRepository.GetAll(
			context.Background(),
			"639c088d06b3b786e7bb9aee",
			1, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, photos)
	})
}

func TestPhotoRepositoryGetOne(t *testing.T) {
	photoRepository := PhotoRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := photoRepository.GetOne(context.Background(), "639c0c8f525aa742fe445bdb")
		assert.Error(t, err)
	})
}

func TestPhotoRepositoryCreate(t *testing.T) {
	photoRepository := PhotoRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := photoRepository.Create(context.Background(), &models.Photo{})
		assert.Error(t, err)
	})
}

func TestPhotoRepositoryDeleteOne(t *testing.T) {
	photoRepository := PhotoRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := photoRepository.DeleteOne(ctxWithCancel, "639c0ef83ca0870a701a5197")
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := photoRepository.DeleteOne(context.Background(), "639c0ef83ca0870a701a5197")
		assert.NoError(t, err)
	})
}

func TestPhotoRepositoryCount(t *testing.T) {
	photoRepository := PhotoRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := photoRepository.Count(context.Background(), "639c0ef83ca0870a701a5197")
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
