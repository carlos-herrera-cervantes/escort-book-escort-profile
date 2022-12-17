package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServiceControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockServiceRepository := mocks.NewMockIServiceRepository(
		controller,
	)
	mockServiceCategoryRepository := mocks.NewMockIServiceCategoryRepository(
		controller,
	)
	serviceController := ServiceController{
		Repository:                mockServiceRepository,
		ServiceCategoryRepository: mockServiceCategoryRepository,
	}

	t.Run("Should return error when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/service?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63964a969884e0ad66b22046")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := serviceController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/service?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63964a969884e0ad66b22046")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Service{}, errors.New("dummy error")).
			Times(1)

		response := serviceController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/service?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63964a969884e0ad66b22046")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Service{}, nil).
			Times(1)
		mockServiceRepository.
			EXPECT().
			Count(gomock.Any(), gomock.Any()).
			Return(0, nil).
			Times(1)

		response := serviceController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestServiceControllerCreate(t *testing.T) {
	controller := gomock.NewController(t)
	mockServiceRepository := mocks.NewMockIServiceRepository(
		controller,
	)
	mockServiceCategoryRepository := mocks.NewMockIServiceCategoryRepository(
		controller,
	)
	serviceController := ServiceController{
		Repository:                mockServiceRepository,
		ServiceCategoryRepository: mockServiceCategoryRepository,
	}

	t.Run("Should return error when category does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/service",
			strings.NewReader(`{
                "serviceCategoryId": "6396d0fb7a44abe3891ca7aa",
                "cost": 100
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6396d0d241ef261c2555b8d0")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ServiceCategory{}, errors.New("dummy error")).
			Times(1)

		response := serviceController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when service is invalid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/service",
			strings.NewReader(`{
                "cost": 0
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6396d0d241ef261c2555b8d0")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ServiceCategory{}, nil).
			Times(1)

		response := serviceController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when service creation fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/service",
			strings.NewReader(`{
                "serviceCategoryId": "6396d0fb7a44abe3891ca7aa",
                "cost": 100
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6396d0d241ef261c2555b8d0")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ServiceCategory{}, nil).
			Times(1)
		mockServiceRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := serviceController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 201 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/service",
			strings.NewReader(`{
                "serviceCategoryId": "6396d0fb7a44abe3891ca7aa",
                "cost": 100
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6396d0d241ef261c2555b8d0")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ServiceCategory{}, nil).
			Times(1)
		mockServiceRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := serviceController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestServiceControllerDeleteOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockServiceRepository := mocks.NewMockIServiceRepository(
		controller,
	)
	mockServiceCategoryRepository := mocks.NewMockIServiceCategoryRepository(
		controller,
	)
	serviceController := ServiceController{
		Repository:                mockServiceRepository,
		ServiceCategoryRepository: mockServiceCategoryRepository,
	}

	t.Run("Should return error when service does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/service/:id",
			nil,
		)
		request.Header.Set("user-id", "6396d0d241ef261c2555b8d0")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Service{}, errors.New("dummy error")).
			Times(1)

		response := serviceController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when service deletion fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/service/:id",
			nil,
		)
		request.Header.Set("user-id", "6396d0d241ef261c2555b8d0")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Service{}, nil).
			Times(1)
		mockServiceRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := serviceController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/service/:id",
			nil,
		)
		request.Header.Set("user-id", "6396d0d241ef261c2555b8d0")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockServiceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.Service{}, nil).
			Times(1)
		mockServiceRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := serviceController.DeleteOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
