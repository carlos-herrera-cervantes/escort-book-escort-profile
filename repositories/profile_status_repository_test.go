package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestProfileStatusRepositoryGetOne(t *testing.T) {
	profileStatusRepository := ProfileStatusRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := profileStatusRepository.GetOne(
			context.Background(),
			"639cbeebfd625a79bfe9c477",
		)
		assert.Error(t, err)
	})
}

func TestProfileStatusRepositoryCreate(t *testing.T) {
	profileStatusRepository := ProfileStatusRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := profileStatusRepository.Create(
			context.Background(),
			&models.ProfileStatus{},
		)
		assert.Error(t, err)
	})
}

func TestProfileStatusRepositoryUpdateOne(t *testing.T) {
	profileStatusRepository := ProfileStatusRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := profileStatusRepository.UpdateOne(
			ctxWithCancel,
			"639cc032f5f77eacd65d1a2f",
			&models.ProfileStatus{},
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := profileStatusRepository.UpdateOne(
			context.Background(),
			"639cc032f5f77eacd65d1a2f",
			&models.ProfileStatus{},
		)
		assert.NoError(t, err)
	})
}
