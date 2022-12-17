package controllers

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories/mocks"
	mockServices "escort-book-escort-profile/services/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPhotoControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockPhotoRepository := mocks.NewMockIPhotoRepository(controller)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	photoController := PhotoController{
		Repository: mockPhotoRepository,
		S3Service:  mockS3Service,
	}

	t.Run("Should return error when pager is invalid", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "-1")
		query.Set("limit", "12")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/photos?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63924ec05cead4fbb48c9216")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		response := photoController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/photos?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63924ec05cead4fbb48c9216")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockPhotoRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Photo{}, errors.New("dummy error")).
			Times(1)

		response := photoController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		query := make(url.Values)
		query.Set("offset", "0")
		query.Set("limit", "10")

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/photos?"+query.Encode(),
			nil,
		)
		request.Header.Set("user-id", "63924ec05cead4fbb48c9216")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockPhotoRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]models.Photo{}, nil).
			Times(1)
		mockPhotoRepository.
			EXPECT().
			Count(gomock.Any(), gomock.Any()).
			Return(0, nil).
			Times(1)

		response := photoController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestPhotoControllerCreate(t *testing.T) {
	controller := gomock.NewController(t)
	mockPhotoRepository := mocks.NewMockIPhotoRepository(controller)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	photoController := PhotoController{
		Repository: mockPhotoRepository,
		S3Service:  mockS3Service,
	}

	t.Run("Should return error when S3 fails", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/photos",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63916dc72171db1919ce0389")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", errors.New("dummy error")).
			Times(1)

		response := photoController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when photo is invalid", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/photos",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)

		response := photoController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when photo creation fails", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/photos",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63916dc72171db1919ce0389")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)
		mockPhotoRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := photoController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return 201 when process succeeds", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/photos",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63916dc72171db1919ce0389")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)
		mockPhotoRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := photoController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestPhotoControllerDeleteOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockPhotoRepository := mocks.NewMockIPhotoRepository(controller)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	photoController := PhotoController{
		Repository: mockPhotoRepository,
		S3Service:  mockS3Service,
	}

	t.Run("Should return error when photo does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/photos/:id",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6392d831fe6a7bf8708cdf0d")

		mockPhotoRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Photo{}, errors.New("dummy error")).
			Times(1)

		response := photoController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when photo deletion fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/photos/:id",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6392d831fe6a7bf8708cdf0d")

		mockPhotoRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Photo{}, nil).
			Times(1)
		mockPhotoRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := photoController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/photos/:id",
			nil,
		)
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("6392d831fe6a7bf8708cdf0d")

		mockPhotoRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Photo{}, nil).
			Times(1)
		mockPhotoRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := photoController.DeleteOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
