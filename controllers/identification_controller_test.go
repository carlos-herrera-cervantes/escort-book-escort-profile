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

func TestIdentificationControllerGetAll(t *testing.T) {
	controller := gomock.NewController(t)
	mockIdentificationRepository := mocks.NewMockIIdentificationRepository(controller)
	mockIdentificationPartRepository := mocks.NewMockIIdentificationPartRepository(
		controller,
	)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	identificationController := IdentificationController{
		Repository:                       mockIdentificationRepository,
		IdentificationCategoryRepository: mockIdentificationPartRepository,
		S3Service:                        mockS3Service,
	}

	t.Run("Should return error when repository fails", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/identifications",
			nil,
		)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockIdentificationRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any()).
			Return([]models.Identification{}, errors.New("dummy error")).
			Times(1)

		response := identificationController.GetAll(c)

		assert.Error(t, response)
	})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		e := echo.New()

		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/escort/profile/identifications",
			nil,
		)
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockIdentificationRepository.
			EXPECT().
			GetAll(gomock.Any(), gomock.Any()).
			Return([]models.Identification{}, nil).
			Times(1)

		response := identificationController.GetAll(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestIdentificationControllerCreate(t *testing.T) {
	controller := gomock.NewController(t)
	mockIdentificationRepository := mocks.NewMockIIdentificationRepository(controller)
	mockIdentificationPartRepository := mocks.NewMockIIdentificationPartRepository(
		controller,
	)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	identificationController := IdentificationController{
		Repository:                       mockIdentificationRepository,
		IdentificationCategoryRepository: mockIdentificationPartRepository,
		S3Service:                        mockS3Service,
	}

	t.Run("Should return error when category does not exists", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		_ = writer.WriteField("identificationPartId", "6392414428a3946dcea53a1a")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/identifications",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockIdentificationPartRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.IdentificationPart{}, errors.New("dummy error")).
			Times(1)

		response := identificationController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when S3 fails", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		_ = writer.WriteField("identificationPartId", "6392414428a3946dcea53a1a")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/identifications",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockIdentificationPartRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.IdentificationPart{}, nil).
			Times(1)
		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", errors.New("dummy error")).
			Times(1)

		response := identificationController.Create(c)

		assert.Error(t, response)
	})

	t.Run("Should return error when identification is invalid", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		_ = writer.WriteField("identificationPartId", "6392414428a3946dcea53a1a")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/identifications",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockIdentificationPartRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.IdentificationPart{}, nil).
			Times(1)
		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)

		response := identificationController.Create(c)

		assert.Error(t, response)
	})

	t.Run(
		"Should return error when identification creation fails",
		func(t *testing.T) {
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
			_ = writer.WriteField("identificationPartId", "6392414428a3946dcea53a1a")
			writer.Close()

			e := echo.New()

			request := httptest.NewRequest(
				http.MethodPost,
				"/api/v1/escort/profile/identifications",
				body,
			)
			request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			request.Header.Set("user-id", "63920ca454b5143e79532d8a")
			recorder := httptest.NewRecorder()

			c := e.NewContext(request, recorder)

			mockIdentificationPartRepository.
				EXPECT().
				GetById(gomock.Any(), gomock.Any()).
				Return(models.IdentificationPart{}, nil).
				Times(1)
			mockS3Service.
				EXPECT().
				Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return("", nil).
				Times(1)
			mockIdentificationRepository.
				EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(errors.New("dummy error")).
				Times(1)

			response := identificationController.Create(c)

			assert.Error(t, response)
		})

	t.Run("Should return 201 when process succeeds", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		_ = writer.WriteField("identificationPartId", "6392414428a3946dcea53a1a")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/escort/profile/identifications",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)

		mockIdentificationPartRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.IdentificationPart{}, nil).
			Times(1)
		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)
		mockIdentificationRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := identificationController.Create(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}

func TestIdentificationControllerUpdateOne(t *testing.T) {
	controller := gomock.NewController(t)
	mockIdentificationRepository := mocks.NewMockIIdentificationRepository(controller)
	mockIdentificationPartRepository := mocks.NewMockIIdentificationPartRepository(
		controller,
	)
	mockS3Service := mockServices.NewMockIS3Service(controller)
	identificationController := IdentificationController{
		Repository:                       mockIdentificationRepository,
		IdentificationCategoryRepository: mockIdentificationPartRepository,
		S3Service:                        mockS3Service,
	}

	t.Run("Should return error when category does not exists", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile/identifications/:id",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("63924ec05cead4fbb48c9216")

		mockIdentificationPartRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.IdentificationPart{}, errors.New("dummy error")).
			Times(1)

		response := identificationController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run(
		"Should return error when identification does not exists",
		func(t *testing.T) {
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
			writer.Close()

			e := echo.New()

			request := httptest.NewRequest(
				http.MethodPatch,
				"/api/v1/escort/profile/identifications/:id",
				body,
			)
			request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			request.Header.Set("user-id", "63920ca454b5143e79532d8a")
			recorder := httptest.NewRecorder()

			c := e.NewContext(request, recorder)
			c.SetParamNames("id")
			c.SetParamValues("63924ec05cead4fbb48c9216")

			mockIdentificationPartRepository.
				EXPECT().
				GetById(gomock.Any(), gomock.Any()).
				Return(models.IdentificationPart{}, nil).
				Times(1)
			mockIdentificationRepository.
				EXPECT().
				GetOne(gomock.Any(), gomock.Any()).
				Return(models.Identification{}, errors.New("dummy error")).
				Times(1)

			response := identificationController.UpdateOne(c)

			assert.Error(t, response)
		})

	t.Run("Should return error when S3 fails", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile/identifications/:id",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("63924ec05cead4fbb48c9216")

		mockIdentificationPartRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.IdentificationPart{}, nil).
			Times(1)
		mockIdentificationRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Identification{}, nil).
			Times(1)
		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", errors.New("dummy error")).
			Times(1)

		response := identificationController.UpdateOne(c)

		assert.Error(t, response)
	})

	t.Run(
		"Should return error when identification update fails",
		func(t *testing.T) {
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
			writer.Close()

			e := echo.New()

			request := httptest.NewRequest(
				http.MethodPatch,
				"/api/v1/escort/profile/identifications/:id",
				body,
			)
			request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			request.Header.Set("user-id", "63920ca454b5143e79532d8a")
			recorder := httptest.NewRecorder()

			c := e.NewContext(request, recorder)
			c.SetParamNames("id")
			c.SetParamValues("63924ec05cead4fbb48c9216")

			mockIdentificationPartRepository.
				EXPECT().
				GetById(gomock.Any(), gomock.Any()).
				Return(models.IdentificationPart{}, nil).
				Times(1)
			mockIdentificationRepository.
				EXPECT().
				GetOne(gomock.Any(), gomock.Any()).
				Return(models.Identification{}, nil).
				Times(1)
			mockS3Service.
				EXPECT().
				Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return("", nil).
				Times(1)
			mockIdentificationRepository.
				EXPECT().
				UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(errors.New("dummy error")).
				Times(1)

			response := identificationController.UpdateOne(c)

			assert.Error(t, response)
		})

	t.Run("Should return 200 when process succeeds", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		_, _ = writer.CreateFormFile("image", "../assets/avatar.png")
		writer.Close()

		e := echo.New()

		request := httptest.NewRequest(
			http.MethodPatch,
			"/api/v1/escort/profile/identifications/:id",
			body,
		)
		request.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
		request.Header.Set("user-id", "63920ca454b5143e79532d8a")
		recorder := httptest.NewRecorder()

		c := e.NewContext(request, recorder)
		c.SetParamNames("id")
		c.SetParamValues("63924ec05cead4fbb48c9216")

		mockIdentificationPartRepository.
			EXPECT().
			GetById(gomock.Any(), gomock.Any()).
			Return(models.IdentificationPart{}, nil).
			Times(1)
		mockIdentificationRepository.
			EXPECT().
			GetOne(gomock.Any(), gomock.Any()).
			Return(models.Identification{}, nil).
			Times(1)
		mockS3Service.
			EXPECT().
			Upload(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", nil).
			Times(1)
		mockIdentificationRepository.
			EXPECT().
			UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		response := identificationController.UpdateOne(c)

		assert.NoError(t, response)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
