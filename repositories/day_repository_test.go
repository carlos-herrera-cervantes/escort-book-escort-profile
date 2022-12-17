package repositories

import (
	"context"
	"testing"

	"escort-book-escort-profile/singleton"

	"github.com/stretchr/testify/assert"
)

func TestDayRepositoryGetAll(t *testing.T) {
	dayRepository := DayRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := dayRepository.GetAll(ctxWithCancel, 1, 10)
		assert.Error(t, err)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		days, err := dayRepository.GetAll(context.Background(), 1, 10)
		assert.Empty(t, days)
		assert.NoError(t, err)
	})
}

func TestDayRepositoryGetOneByName(t *testing.T) {
	dayRepository := DayRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := dayRepository.GetOneByName(context.Background(), "Monday")
		assert.Error(t, err)
	})
}

func TestDayRepositoryGetById(t *testing.T) {
	dayRepository := DayRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := dayRepository.GetById(
			context.Background(),
			"639abb4e9527be53adb77288",
		)
		assert.Error(t, err)
	})
}

func TestDayRepositoryCount(t *testing.T) {
	dayRepository := DayRepository{
		Data: singleton.NewPostgresClient(),
	}

	t.Run("Should return 0", func(t *testing.T) {
		counter, err := dayRepository.Count(context.Background())
		assert.Zero(t, counter)
		assert.NoError(t, err)
	})
}
