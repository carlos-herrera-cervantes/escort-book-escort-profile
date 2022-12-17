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

func TestAttentionSiteControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockAttentionSiteRepository := mocks.NewMockIAttentionSiteRepository(controller)
	mockAttentionSiteCategoryRepository := mocks.NewMockIAttentionSiteCategoryRepository(controller)
	attentionSiteController := AttentionSiteController{
		Repository:                      mockAttentionSiteRepository,
		AttentionSiteCategoryRepository: mockAttentionSiteCategoryRepository,
	}

	t.Run("Should return bad request when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/attention-sites?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := attentionSiteController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return internal server error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/attention-sites?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAttentionSiteRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.AttentionSiteDetailed{}, errors.New("dummy error")).
			Times(1)
		mockAttentionSiteRepository.
			EXPECT().
			Count(gomock.Any(), gomock.Any()).
			Return(0, nil).
			Times(1)

		response := attentionSiteController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/attention-sites?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAttentionSiteRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.AttentionSiteDetailed{}, nil).
			Times(1)
		mockAttentionSiteRepository.
			EXPECT().
			Count(gomock.Any(), gomock.Any()).
			Return(0, nil).
			Times(1)

		response := attentionSiteController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAttentionSiteControllerGetById(t *testing.T) {
	controller := gomock.NewController(t)
	mockAttentionSiteRepository := mocks.NewMockIAttentionSiteRepository(controller)
	mockAttentionSiteCategoryRepository := mocks.NewMockIAttentionSiteCategoryRepository(controller)
	attentionSiteController := AttentionSiteController{
		Repository:                      mockAttentionSiteRepository,
		AttentionSiteCategoryRepository: mockAttentionSiteCategoryRepository,
	}

	t.Run("Should return bad request when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/:id/profile/attention-sites?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("639152a716d1b6fca0c161f7")

		response := attentionSiteController.GetById(c)

		assert.Error(t, response)
	})

	t.Run("Should return internal server error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/:id/profile/attention-sites?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("639152a716d1b6fca0c161f7")

		mockAttentionSiteRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.AttentionSiteDetailed{}, errors.New("dummy error")).
			Times(1)
		mockAttentionSiteRepository.
			EXPECT().
			Count(gomock.Any(), gomock.Any()).
			Return(0, nil).
			Times(1)

		response := attentionSiteController.GetById(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/:id/profile/attention-sites?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("639152a716d1b6fca0c161f7")

		mockAttentionSiteRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.AttentionSiteDetailed{}, nil).
			Times(1)
		mockAttentionSiteRepository.
			EXPECT().
			Count(gomock.Any(), gomock.Any()).
			Return(0, nil).
			Times(1)

		response := attentionSiteController.GetById(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAttentionSiteControllerCreate(t *testing.T) {
	controller := gomock.NewController(t)
	mockAttentionSiteRepository := mocks.NewMockIAttentionSiteRepository(controller)
	mockAttentionSiteCategoryRepository := mocks.NewMockIAttentionSiteCategoryRepository(controller)
	attentionSiteController := AttentionSiteController{
		Repository:                      mockAttentionSiteRepository,
		AttentionSiteCategoryRepository: mockAttentionSiteCategoryRepository,
	}

	t.Run("Should return error when attention site category does not exists", func(t *testing.T) {
		e := echo.New()
		body := `{"attentionSiteCategoryId": "63915d29fea7c3da958e20c2"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/attention-sites",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "63915d9be04ff7e3da92ffaa")

		mockAttentionSiteCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.AttentionSiteCategory{}, errors.New("dummy error")).
			Times(1)

		response := attentionSiteController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when attention site is invalid", func(t *testing.T) {
		e := echo.New()
		body := `{}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/attention-sites",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "63915d9be04ff7e3da92ffaa")

		mockAttentionSiteCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.AttentionSiteCategory{}, nil).
			Times(1)

		response := attentionSiteController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when attention site creation fails", func(t *testing.T) {
		e := echo.New()
		body := `{"attentionSiteCategoryId": "63915d29fea7c3da958e20c2"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/attention-sites",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "63915d9be04ff7e3da92ffaa")

		mockAttentionSiteCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.AttentionSiteCategory{}, nil).
			Times(1)
		mockAttentionSiteRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := attentionSiteController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 201 when process succeeds", func(t *testing.T) {
		e := echo.New()
		body := `{"attentionSiteCategoryId": "63915d29fea7c3da958e20c2"}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/attention-sites",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.Request().Header.Set("user-id", "63915d9be04ff7e3da92ffaa")

		mockAttentionSiteCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.AttentionSiteCategory{}, nil).
			Times(1)
		mockAttentionSiteRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := attentionSiteController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestAttentionSiteControllerDeleteOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockAttentionSiteRepository := mocks.NewMockIAttentionSiteRepository(controller)
	mockAttentionSiteCategoryRepository := mocks.NewMockIAttentionSiteCategoryRepository(controller)
	attentionSiteController := AttentionSiteController{
		Repository:                      mockAttentionSiteRepository,
		AttentionSiteCategoryRepository: mockAttentionSiteCategoryRepository,
	}

	t.Run("Should return error when attention site does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/attention-sites/:id",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("639152a716d1b6fca0c161f7")

		mockAttentionSiteRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.AttentionSiteDetailed{}, errors.New("dummy error")).
			Times(1)

		response := attentionSiteController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when attention site deletion fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/attention-sites/:id",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("639152a716d1b6fca0c161f7")

		mockAttentionSiteRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.AttentionSiteDetailed{}, nil).
			Times(1)
		mockAttentionSiteRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := attentionSiteController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/attention-sites/:id",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("639152a716d1b6fca0c161f7")

		mockAttentionSiteRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.AttentionSiteDetailed{}, nil).
			Times(1)
		mockAttentionSiteRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := attentionSiteController.DeleteOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
