package controllers

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories/mocks"
	mockServices "escort-book-escort-profile/services/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAvatarControllerGetOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(controller)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	avatarController := AvatarController{
		Repository: mockAvatarRepository,
		S3Service:  mockS3Service,
	}

	t.Run("Should return error when avatar does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/avatar",
			nil,
		)
		request.Header.Set("user-id", "63916dc72171db1919ce0389")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAvatarRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Avatar{}, errors.New("dummy error")).
			Times(1)

		response := avatarController.GetOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/avatar",
			nil,
		)
		request.Header.Set("user-id", "63916dc72171db1919ce0389")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAvatarRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Avatar{}, nil).
			Times(1)

		response := avatarController.GetOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestAvatarControllerUpsert(t *testing.T) {
	controller := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(controller)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	avatarController := AvatarController{
		Repository: mockAvatarRepository,
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
			"/api/v1/escort/profile/avatar",
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

		response := avatarController.Upsert(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when avatar is invalid", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/avatar",
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

		response := avatarController.Upsert(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when avatar creation fails", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/avatar",
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
		mockAvatarRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Avatar{}, errors.New("dummy error")).
			Times(1)
		mockAvatarRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := avatarController.Upsert(c)

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
			"/api/v1/escort/profile/avatar",
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
		mockAvatarRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Avatar{}, errors.New("dummy error")).
			Times(1)
		mockAvatarRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := avatarController.Upsert(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestAvatarControllerDeleteOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockAvatarRepository := mocks.NewMockIAvatarRepository(controller)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	avatarController := AvatarController{
		Repository: mockAvatarRepository,
		S3Service:  mockS3Service,
	}

	t.Run("Should return error when avatar does not exists", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/avatar",
			nil,
		)
		request.Header.Set("user-id", "63916dc72171db1919ce0389")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAvatarRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Avatar{}, errors.New("dummy error")).
			Times(1)

		response := avatarController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when avatar removal fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/avatar",
			nil,
		)
		request.Header.Set("user-id", "63916dc72171db1919ce0389")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAvatarRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Avatar{}, nil).
			Times(1)
		mockAvatarRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		response := avatarController.DeleteOne(c)

		assert.Error(t, response)
	})

	t.Run("Should return 204 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodDelete,
			"/api/v1/escort/profile/avatar",
			nil,
		)
		request.Header.Set("user-id", "63916dc72171db1919ce0389")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockAvatarRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Avatar{}, nil).
			Times(1)
		mockAvatarRepository.
			EXPECT().
			DeleteOne(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := avatarController.DeleteOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}
