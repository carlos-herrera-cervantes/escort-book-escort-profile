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

func TestScheduleControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockScheduleRepository := mocks.NewMockIScheduleRepository(controller)
	mockDayRepository := mocks.NewMockIDayRepository(controller)
	scheduleController := ScheduleController{
		Repository:    mockScheduleRepository,
		DayRepository: mockDayRepository,
	}

	t.Run("Should return error when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/schedules?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "6394cb5ade78cbc364a6f4c6")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := scheduleController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/schedules?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "6394cb5ade78cbc364a6f4c6")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockScheduleRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Schedule{}, errors.New("dummy error")).
			Times(1)

		response := scheduleController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/schedules?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "6394cb5ade78cbc364a6f4c6")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockScheduleRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Schedule{}, nil).
			Times(1)
		mockScheduleRepository.
			EXPECT().
			Count(gomock.Any(), gomock.Any()).
			Return(0, nil)

		response := scheduleController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestScheduleControllerCreate(t *testing.T) {
	controller := gomock.NewController(t)
	mockScheduleRepository := mocks.NewMockIScheduleRepository(controller)
	mockDayRepository := mocks.NewMockIDayRepository(controller)
	scheduleController := ScheduleController{
		Repository:    mockScheduleRepository,
		DayRepository: mockDayRepository,
	}

	t.Run("Should return error when day does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1//escort/profile/schedules",
			strings.NewReader(`{
                "from": "2022-01-01",
                "to": "2022-12-31",
                "dayId": "63963b0fc438aadd5efe167c"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63963ab9f17f0967331c7c9e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockDayRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.Day{}, errors.New("dummy error")).
			Times(1)

		response := scheduleController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when schedule is invalid", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1//escort/profile/schedules",
			strings.NewReader(`{
                "from": "2022-01-01",
                "to": "2022-12-31",
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63963ab9f17f0967331c7c9e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockDayRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.Day{}, nil).
			Times(1)

		response := scheduleController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when schedule creation fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1//escort/profile/schedules",
			strings.NewReader(`{
                "from": "2022-01-01",
                "to": "2022-12-31",
                "dayId": "63963b0fc438aadd5efe167c"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63963ab9f17f0967331c7c9e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockDayRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.Day{}, nil).
			Times(1)
		mockScheduleRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := scheduleController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 201 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1//escort/profile/schedules",
			strings.NewReader(`{
                "from": "2022-01-01",
                "to": "2022-12-31",
                "dayId": "63963b0fc438aadd5efe167c"
            }`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63963ab9f17f0967331c7c9e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockDayRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.Day{}, nil).
			Times(1)
		mockScheduleRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := scheduleController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestScheduleControllerDeleteOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockScheduleRepository := mocks.NewMockIScheduleRepository(controller)
	mockDayRepository := mocks.NewMockIDayRepository(controller)
	scheduleController := ScheduleController{
		Repository:    mockScheduleRepository,
		DayRepository: mockDayRepository,
	}

	t.Run("Should return error when schedule does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/schedules/:id",
			nil,
		)
		request.Header.Set("user-id", "63963ab9f17f0967331c7c9e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6396464e81053081dc137365")

		mockScheduleRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Schedule{}, errors.New("dummy error")).
			Times(1)

		response := scheduleController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when schedule deletion fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/schedules/:id",
			nil,
		)
		request.Header.Set("user-id", "63963ab9f17f0967331c7c9e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6396464e81053081dc137365")

		mockScheduleRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Schedule{}, nil).
			Times(1)
		mockScheduleRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := scheduleController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/schedules/:id",
			nil,
		)
		request.Header.Set("user-id", "63963ab9f17f0967331c7c9e")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6396464e81053081dc137365")

		mockScheduleRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Schedule{}, nil).
			Times(1)
		mockScheduleRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := scheduleController.DeleteOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
