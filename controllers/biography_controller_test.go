package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBiographyControllerGetOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockBiographyRepository := mocks.NewMockIBiographyRepository(controller)
	biographyController := BiographyController{
		Repository: mockBiographyRepository,
	}

	t.Run("Should return error when biography does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/biography",
			nil,
		)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Biography{}, errors.New("dummy error")).
			Times(1)

		response := biographyController.GetOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/biography",
			nil,
		)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Biography{}, nil).
			Times(1)

		response := biographyController.GetOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestBiographyControllerCreate(t *testing.T) {
	controller := gomock.NewController(t)
	mockBiographyRepository := mocks.NewMockIBiographyRepository(controller)
	biographyController := BiographyController{
		Repository: mockBiographyRepository,
	}

	t.Run("Should return error when biography is invalid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/biography",
			strings.NewReader("{}"),
		)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := biographyController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when biography creation fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/biography",
			strings.NewReader(`{"description": "Test description"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := biographyController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 201 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/biography",
			strings.NewReader(`{"description": "Test description"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := biographyController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestBiographyControllerUpdateOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockBiographyRepository := mocks.NewMockIBiographyRepository(controller)
	biographyController := BiographyController{
		Repository: mockBiographyRepository,
	}

	t.Run("Should return error when biography is invalid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/v1/escort/profile/biography",
			strings.NewReader(`{}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := biographyController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when biography does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/v1/escort/profile/biography",
			strings.NewReader(`{"description": "Test description"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Biography{}, errors.New("dummy error")).
			Times(1)

		response := biographyController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when biography update fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/v1/escort/profile/biography",
			strings.NewReader(`{"description": "Test description"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Biography{}, nil).
			Times(1)
		mockBiographyRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := biographyController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/v1/escort/profile/biography",
			strings.NewReader(`{"description": "Test description"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Biography{}, nil).
			Times(1)
		mockBiographyRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := biographyController.UpdateOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestBiographyControllerDeleteOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockBiographyRepository := mocks.NewMockIBiographyRepository(controller)
	biographyController := BiographyController{
		Repository: mockBiographyRepository,
	}

	t.Run("Should return error when biography does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/biography",
			nil,
		)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Biography{}, errors.New("dummy error")).
			Times(1)

		response := biographyController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when biography removing fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/biography",
			nil,
		)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Biography{}, nil).
			Times(1)
		mockBiographyRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := biographyController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/biography",
			nil,
		)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockBiographyRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Biography{}, nil).
			Times(1)
		mockBiographyRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := biographyController.DeleteOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
