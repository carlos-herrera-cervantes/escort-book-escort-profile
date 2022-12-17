package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories/mocks"
	mockServices "escort-book-escort-profile/services/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProfileStatusControllerGetByExternal(t *testing.T) {
	controller := gomock.NewController(t)
	mockProfileStatusRepository := mocks.NewMockIProfileStatusRepository(controller)
	mockProfileStatusCategoryRepository := mocks.NewMockIProfileStatusCategoryRepository(
		controller,
	)
	mockKafkaService := mockServices.NewMockIKafkaService(controller)
	profileStatusController := ProfileStatusController{
		Repository:                      mockProfileStatusRepository,
		ProfileStatusCategoryRepository: mockProfileStatusCategoryRepository,
		KafkaService:                    mockKafkaService,
	}

	t.Run("Should return error when status does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/:id/profile/status",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6394c8c3960b6e38c1865b6e")

		mockProfileStatusRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.All()).
			Return(models.ProfileStatus{}, errors.New("dummy error")).
			Times(1)

		response := profileStatusController.GetByExternal(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/:id/profile/status",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6394c8c3960b6e38c1865b6e")

		mockProfileStatusRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.All()).
			Return(models.ProfileStatus{}, nil).
			Times(1)

		response := profileStatusController.GetByExternal(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestProfileStatusControllerUpdateOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockProfileStatusRepository := mocks.NewMockIProfileStatusRepository(controller)
	mockProfileStatusCategoryRepository := mocks.NewMockIProfileStatusCategoryRepository(
		controller,
	)
	mockKafkaService := mockServices.NewMockIKafkaService(controller)
	profileStatusController := ProfileStatusController{
		Repository:                      mockProfileStatusRepository,
		ProfileStatusCategoryRepository: mockProfileStatusCategoryRepository,
		KafkaService:                    mockKafkaService,
	}

	t.Run("Should return error when category does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile/status",
			strings.NewReader(`{
                "profileStatusCategoryId": "6394cca01cc92721d326f566"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6394cb5ade78cbc364a6f4c6")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileStatusCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{}, errors.New("dummy error")).
			Times(1)

		response := profileStatusController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when operation is invalid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile/status",
			strings.NewReader(`{
                "profileStatusCategoryId": "6394cca01cc92721d326f566"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6394cb5ade78cbc364a6f4c6")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileStatusCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{
				Name: "Locked",
			}, nil).
			Times(1)

		response := profileStatusController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error whe status does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile/status",
			strings.NewReader(`{
                "profileStatusCategoryId": "6394cca01cc92721d326f566"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6394cb5ade78cbc364a6f4c6")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileStatusCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{
				Name: "Deactivated",
			}, nil).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatus{}, errors.New("dummy error")).
			Times(1)

		response := profileStatusController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when status update fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile/status",
			strings.NewReader(`{
                "profileStatusCategoryId": "6394cca01cc92721d326f566"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6394cb5ade78cbc364a6f4c6")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileStatusCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{
				Name: "Deactivated",
			}, nil).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatus{}, nil).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := profileStatusController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile/status",
			strings.NewReader(`{
                "profileStatusCategoryId": "6394cca01cc92721d326f566"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "6394cb5ade78cbc364a6f4c6")
		request.Header.Set("user-type", "Customer")
		request.Header.Set("user-email", "test.customer@example.com")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileStatusCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatusCategory{
				Name: "Deleted",
			}, nil).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.ProfileStatus{}, nil).
			Times(1)
		mockProfileStatusRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)
		mockKafkaService.
			EXPECT().
			SendMessage(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(2)

		response := profileStatusController.UpdateOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
