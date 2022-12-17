package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestProfileStatusCategoryRepositoryGetAll(t *testing.T) {
	profileStatusCategoryRepository := ProfileStatusCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := profileStatusCategoryRepository.GetAll(ctxWithCancel, 1, 10)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		status, err := profileStatusCategoryRepository.GetAll(
			context.Background(),
			1, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, status)
	})
}

func TestProfileStatusCategoryRepositoryGetOneByName(t *testing.T) {
	profileStatusCategoryRepository := ProfileStatusCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := profileStatusCategoryRepository.GetOneByName(
			context.Background(),
			"dummy",
		)
		assert.Error(t, err)
	})
}

func TestProfileStatusCategoryRepositoryCount(t *testing.T) {
	profileStatusCategoryRepository := ProfileStatusCategoryRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := profileStatusCategoryRepository.Count(context.Background())
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
