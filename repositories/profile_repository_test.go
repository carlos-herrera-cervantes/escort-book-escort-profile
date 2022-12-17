package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestProfileRepositoryGetAll(t *testing.T) {
	profileRepository := ProfileRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := profileRepository.GetAll(ctxWithCancel, 1, 10)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		profiles, err := profileRepository.GetAll(context.Background(), 1, 10)
		assert.NoError(t, err)
		assert.Empty(t, profiles)
	})
}

func TestProfileRepositoryGetOne(t *testing.T) {
	profileRepository := ProfileRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := profileRepository.GetOne(
			context.Background(),
			"639cad2a97924938a03336f1",
		)
		assert.Error(t, err)
	})
}

func TestProfileRepositoryCreate(t *testing.T) {
	profileRepository := ProfileRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := profileRepository.Create(context.Background(), &models.Profile{})
		assert.Error(t, err)
	})
}

func TestProfileRepositoryUpdateOne(t *testing.T) {
	profileRepository := ProfileRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := profileRepository.UpdateOne(
			context.Background(),
			"639cad2a97924938a03336f1",
			&models.Profile{},
		)
		assert.Error(t, err)
	})
}

func TestProfileRepositoryDeleteOne(t *testing.T) {
	profileRepository := ProfileRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := profileRepository.DeleteOne(
			ctxWithCancel,
			"639cad2a97924938a03336f1",
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := profileRepository.DeleteOne(
			context.Background(),
			"639cad2a97924938a03336f1",
		)
		assert.NoError(t, err)
	})
}

func TestProfileRepositoryCount(t *testing.T) {
	profileRepository := ProfileRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := profileRepository.Count(context.Background())
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
