package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestBiographyRepositoryGetOne(t *testing.T) {
	biographyRepository := BiographyRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := biographyRepository.GetOne(
			context.Background(),
			"6398e9f05299063cb696d08d",
		)
		assert.Error(t, err)
	})
}

func TestBiographyRepositoryCreate(t *testing.T) {
	biographyRepository := BiographyRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := biographyRepository.Create(
			context.Background(),
			&models.Biography{},
		)
		assert.Error(t, err)
	})
}

func TestBiographyRepositoryUpdateOne(t *testing.T) {
	biographyRepository := BiographyRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := biographyRepository.UpdateOne(
			ctxWithCancel,
			"639aa30b01e58d2e9ee70729",
			&models.Biography{},
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := biographyRepository.UpdateOne(
			context.Background(),
			"639aa30b01e58d2e9ee70729",
			&models.Biography{},
		)
		assert.NoError(t, err)
	})
}

func TestBiographyRepositoryDeleteOne(t *testing.T) {
	biographyRepository := BiographyRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := biographyRepository.DeleteOne(
			ctxWithCancel,
			"639aa90aef4edcb2738fc787",
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := biographyRepository.DeleteOne(
			context.Background(),
			"639aa90aef4edcb2738fc787",
		)
		assert.NoError(t, err)
	})
}
