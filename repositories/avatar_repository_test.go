package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestAvatarRepositoryGetOne(t *testing.T) {
	avatarRepository := AvatarRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := avatarRepository.GetOne(
			context.Background(),
			"6398db456fa88a545f2dc24d",
		)
		assert.Error(t, err)
	})
}

func TestAvatarRepositoryCreate(t *testing.T) {
	avatarRepository := AvatarRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := avatarRepository.Create(
			context.Background(),
			&models.Avatar{},
		)
		assert.Error(t, err)
	})
}

func TestAvatarRepositoryUpdateOne(t *testing.T) {
	avatarRepository := AvatarRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := avatarRepository.UpdateOne(
			ctxWithCancel,
			"6398db456fa88a545f2dc24d",
			&models.Avatar{},
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := avatarRepository.UpdateOne(
			context.Background(),
			"6398db456fa88a545f2dc24d",
			&models.Avatar{},
		)
		assert.NoError(t, err)
	})
}
func TestAvatarRepositoryDeleteOne(t *testing.T) {
	avatarRepository := AvatarRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := avatarRepository.DeleteOne(
			ctxWithCancel,
			"6398db456fa88a545f2dc24d",
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := avatarRepository.DeleteOne(
			context.Background(),
			"6398db456fa88a545f2dc24d",
		)
		assert.NoError(t, err)
	})
}
