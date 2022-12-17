package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIdentificationPartControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockIdentificationPartRepository := mocks.NewMockIIdentificationPartRepository(
		controller,
	)
	identificationPartController := IdentificationPartController{
		Repository: mockIdentificationPartRepository,
	}

	t.Run("Should return error when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/identification-parts?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := identificationPartController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/identification-parts?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockIdentificationPartRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.IdentificationPart{}, errors.New("dummy error")).
			Times(1)
		mockIdentificationPartRepository.
			EXPECT().
			Count(gomock.Any()).
			Return(0, nil).
			Times(1)

		response := identificationPartController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/identification-parts?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockIdentificationPartRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.IdentificationPart{}, nil).
			Times(1)
		mockIdentificationPartRepository.
			EXPECT().
			Count(gomock.Any()).
			Return(0, nil).
			Times(1)

		response := identificationPartController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
