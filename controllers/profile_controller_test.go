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
	mockServices "escort-book-escort-profile/services/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProfileControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(controller)
	mockEmitterService := mockServices.NewMockIEmitterService(controller)
	mockNationalityRepository := mocks.NewMockINationalityRepository(
		controller,
	)
	profileController := ProfileController{
		Repository:            mockProfileRepository,
		Emitter:               mockEmitterService,
		NationalityRepository: mockNationalityRepository,
	}

	t.Run("Should return error when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escorts?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := profileController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escorts?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Profile{}, errors.New("dummy error")).
			Times(1)

		response := profileController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escorts?"+query.Encode(),
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Profile{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			Count(gomock.Any()).
			Return(0, nil).
			Times(1)

		response := profileController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestProfileControllerGetOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(controller)
	mockEmitterService := mockServices.NewMockIEmitterService(controller)
	mockNationalityRepository := mocks.NewMockINationalityRepository(
		controller,
	)
	profileController := ProfileController{
		Repository:            mockProfileRepository,
		Emitter:               mockEmitterService,
		NationalityRepository: mockNationalityRepository,
	}

	t.Run("Should return error when profile does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile",
			nil,
		)
		request.Header.Set("user-id", "6393fa9e8d41cf7ce04a78f2")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, errors.New("dummy error")).
			Times(1)

		response := profileController.GetOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile",
			nil,
		)
		request.Header.Set("user-id", "6393fa9e8d41cf7ce04a78f2")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(1)

		response := profileController.GetOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestProfileControllerCreate(t *testing.T) {
	controller := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(controller)
	mockEmitterService := mockServices.NewMockIEmitterService(controller)
	mockNationalityRepository := mocks.NewMockINationalityRepository(
		controller,
	)
	profileController := ProfileController{
		Repository:            mockProfileRepository,
		Emitter:               mockEmitterService,
		NationalityRepository: mockNationalityRepository,
	}

	t.Run("Should return error when nationality does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile",
			strings.NewReader(`{
	            "email": "test.customer@example.com",
	            "phoneNumber": "12345",
	            "gender": "Male"
	        }`),
		)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockNationalityRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.Nationality{}, errors.New("dummy error")).
			Times(1)

		response := profileController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when profile is invalid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile",
			strings.NewReader(`{
	            "email": "test.customer@example.com",
	            "gender": "Male"
	        }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockNationalityRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.Nationality{}, nil).
			Times(1)

		response := profileController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when profile creation fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile",
			strings.NewReader(`{
                "email": "test.customer@example.com",
                "phoneNumber": "12345",
                "gender": "Male"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockNationalityRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.Nationality{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := profileController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 201 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile",
			strings.NewReader(`{
                "email": "test.customer@example.com",
                "phoneNumber": "12345",
                "gender": "Male"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockNationalityRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.Nationality{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)
		mockEmitterService.
			EXPECT().
			Emit(gomock.Any(), gomock.Any()).
			Times(1)

		response := profileController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestProfileControllerUpdateOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(controller)
	mockEmitterService := mockServices.NewMockIEmitterService(controller)
	mockNationalityRepository := mocks.NewMockINationalityRepository(
		controller,
	)
	profileController := ProfileController{
		Repository:            mockProfileRepository,
		Emitter:               mockEmitterService,
		NationalityRepository: mockNationalityRepository,
	}

	t.Run("Should return error when profile does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile",
			strings.NewReader(`{"firstName": "Kamsia"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, errors.New("dummy")).
			Times(1)

		response := profileController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when profile update fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile",
			strings.NewReader(`{"firstName": "Kamsia"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := profileController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile",
			strings.NewReader(`{"firstName": "Kamsia"}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := profileController.UpdateOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestProfileControllerDeleteOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockProfileRepository := mocks.NewMockIProfileRepository(controller)
	mockEmitterService := mockServices.NewMockIEmitterService(controller)
	mockNationalityRepository := mocks.NewMockINationalityRepository(
		controller,
	)
	profileController := ProfileController{
		Repository:            mockProfileRepository,
		Emitter:               mockEmitterService,
		NationalityRepository: mockNationalityRepository,
	}

	t.Run("Should return error when profile does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile",
			nil,
		)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, errors.New("dummy error")).
			Times(1)

		response := profileController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when profile deletion fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile",
			nil,
		)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := profileController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile",
			nil,
		)
		request.Header.Set("user-id", "639406734f4556707f05f97e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockProfileRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Profile{}, nil).
			Times(1)
		mockProfileRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := profileController.DeleteOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
