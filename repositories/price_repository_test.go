package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestPriceRepositoryGetAll(t *testing.T) {
	priceRepository := PriceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := priceRepository.GetAll(ctxWithCancel, "639c0ef83ca0870a701a5197", 1, 10)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		prices, err := priceRepository.GetAll(
			context.Background(),
			"639c0ef83ca0870a701a5197",
			1, 10,
		)
		assert.NoError(t, err)
		assert.Empty(t, prices)
	})
}

func TestPriceRepositoryGetOne(t *testing.T) {
	priceRepository := PriceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := priceRepository.GetOne(context.Background(), "639c13efea5c6e317b87b645")
		assert.Error(t, err)
	})
}

func TestPriceRepositoryCreate(t *testing.T) {
	priceRepository := PriceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := priceRepository.Create(context.Background(), &models.Price{})
		assert.Error(t, err)
	})
}

func TestPriceRepositoryUpdateOne(t *testing.T) {
	priceRepository := PriceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := priceRepository.UpdateOne(
			ctxWithCancel,
			"639c13efea5c6e317b87b645",
			&models.Price{},
		)
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := priceRepository.UpdateOne(
			context.Background(),
			"639c13efea5c6e317b87b645",
			&models.Price{},
		)
		assert.NoError(t, err)
	})
}

func TestPriceRepositoryDeleteOne(t *testing.T) {
	priceRepository := PriceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := priceRepository.DeleteOne(ctxWithCancel, "639c13efea5c6e317b87b645")
		assert.Error(t, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		err := priceRepository.DeleteOne(context.Background(), "639c13efea5c6e317b87b645")
		assert.NoError(t, err)
	})
}

func TestPriceRepositoryCount(t *testing.T) {
	priceRepository := PriceRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := priceRepository.Count(
			context.Background(),
			"639c13efea5c6e317b87b645",
		)
		assert.NoError(t, err)
		assert.Zero(t, counter)
	})
}
