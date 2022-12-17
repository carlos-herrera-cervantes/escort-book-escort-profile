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

func TestAttentionSiteCategoryControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockAttentionSiteCategoryRepository := mocks.NewMockIAttentionSiteCategoryRepository(controller)
	attentionSiteCategoryController := AttentionSiteCategoryController{
		Repository: mockAttentionSiteCategoryRepository,
	}
	
	t.Run("Should return bad request when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/attention-site-categories?" + query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := attentionSiteCategoryController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return internal server error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/attention-site-categories?" + query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAttentionSiteCategoryRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.AttentionSiteCategory{}, errors.New("dummy error")).
			Times(1)
		mockAttentionSiteCategoryRepository.
			EXPECT().
			Count(gomock.Any()).
			Return(0, nil).
			Times(1)

		response := attentionSiteCategoryController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/attention-site-categories?" + query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAttentionSiteCategoryRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.AttentionSiteCategory{}, nil).
			Times(1)
		mockAttentionSiteCategoryRepository.
			EXPECT().
			Count(gomock.Any()).
			Return(0, nil).
			Times(1)

		response := attentionSiteCategoryController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
