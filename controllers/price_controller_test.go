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

func TestPriceControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockPriceRepository := mocks.NewMockIPriceRepository(controller)
	mockTimeCategoryRepository := mocks.NewMockITimeCategoryRepository(
		controller,
	)
	priceController := PriceController{
		Repository:             mockPriceRepository,
		TimeCategoryRepository: mockTimeCategoryRepository,
	}

	t.Run("Should return error when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/prices?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := priceController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/prices?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockPriceRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Price{}, errors.New("dummy error")).
			Times(1)

		response := priceController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/prices?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockPriceRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Price{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			Count(gomock.Any(), gomock.Any()).
			Return(0, nil).
			Times(1)

		response := priceController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestPriceControllerCreate(t *testing.T) {
	controller := gomock.NewController(t)
	mockPriceRepository := mocks.NewMockIPriceRepository(controller)
	mockTimeCategoryRepository := mocks.NewMockITimeCategoryRepository(
		controller,
	)
	priceController := PriceController{
		Repository:             mockPriceRepository,
		TimeCategoryRepository: mockTimeCategoryRepository,
	}

	t.Run("Should return error when category does not exists", func(t *testing.T) {
		e := echo.New()
		body := `{"cost":0,"timeCategoryId":"63938a7d5fee5cd1b6a709e2","quantity":0}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/prices",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockTimeCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.TimeCategory{}, errors.New("dummy error")).
			Times(1)

		response := priceController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when price is invalid", func(t *testing.T) {
		e := echo.New()
		body := `{"timeCategoryId":"63938a7d5fee5cd1b6a709e2","quantity":0}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/prices",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockTimeCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.TimeCategory{}, nil).
			Times(1)

		response := priceController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when price creation fails", func(t *testing.T) {
		e := echo.New()
		body := `{"quantity":1,"timeCategoryId":"63938a7d5fee5cd1b6a709e2","cost":10}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/prices",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockTimeCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.TimeCategory{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := priceController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 201 when process succeeds", func(t *testing.T) {
		e := echo.New()
		body := `{"quantity":1,"timeCategoryId":"63938a7d5fee5cd1b6a709e2","cost":10}`

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/prices",
			strings.NewReader(body),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockTimeCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.TimeCategory{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := priceController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestPriceControllerUpdateOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockPriceRepository := mocks.NewMockIPriceRepository(controller)
	mockTimeCategoryRepository := mocks.NewMockITimeCategoryRepository(
		controller,
	)
	priceController := PriceController{
		Repository:             mockPriceRepository,
		TimeCategoryRepository: mockTimeCategoryRepository,
	}

	t.Run("Should return error when category does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/v1/escort/profile/prices/:id",
			strings.NewReader(`{"cost": 100}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6393a352b27e5d8a955353ae")

		mockTimeCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.TimeCategory{}, errors.New("dummy error")).
			Times(1)

		response := priceController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when price does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/v1/escort/profile/prices/:id",
			strings.NewReader(`{"cost": 100}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6393a352b27e5d8a955353ae")

		mockTimeCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.TimeCategory{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Price{}, errors.New("dummy error")).
			Times(1)

		response := priceController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when price update fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/v1/escort/profile/prices/:id",
			strings.NewReader(`{"cost": 100}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6393a352b27e5d8a955353ae")

		mockTimeCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.TimeCategory{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Price{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := priceController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPut,
			"/api/v1/escort/profile/prices/:id",
			strings.NewReader(`{"cost": 100}`),
		)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6393a352b27e5d8a955353ae")

		mockTimeCategoryRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.TimeCategory{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Price{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := priceController.UpdateOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestPriceControllerDeleteOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockPriceRepository := mocks.NewMockIPriceRepository(controller)
	mockTimeCategoryRepository := mocks.NewMockITimeCategoryRepository(
		controller,
	)
	priceController := PriceController{
		Repository:             mockPriceRepository,
		TimeCategoryRepository: mockTimeCategoryRepository,
	}

	t.Run("Should return error when price does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/prices/:id",
			nil,
		)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6393a352b27e5d8a955353ae")

		mockPriceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Price{}, errors.New("dummy error")).
			Times(1)

		response := priceController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when price deletion fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/prices/:id",
			nil,
		)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6393a352b27e5d8a955353ae")

		mockPriceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Price{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := priceController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/prices/:id",
			nil,
		)
		request.Header.Set("user-id", "63938680374cf000ea97ee3a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6393a352b27e5d8a955353ae")

		mockPriceRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Price{}, nil).
			Times(1)
		mockPriceRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := priceController.DeleteOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
