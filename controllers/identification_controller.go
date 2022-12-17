package controllers

import (
	"fmt"
	"net/http"

	"escort-book-escort-profile/config"
	"escort-book-escort-profile/constants"
	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"

	"github.com/labstack/echo/v4"
)

type IdentificationController struct {
	Repository                       repositories.IIdentificationRepository
	S3Service                        services.IS3Service
	IdentificationCategoryRepository repositories.IIdentificationPartRepository
}

func (h *IdentificationController) GetByExternal(c echo.Context) (err error) {
	id := c.Param("id")
	identifications, err := h.Repository.GetAll(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	for index := range identifications {
		identifications[index].Path = fmt.Sprintf(
			"%s/%s/%s",
			config.InitS3().Endpoint,
			config.InitS3().Buckets.EscortProfile,
			identifications[index].Path,
		)
	}

	return c.JSON(http.StatusOK, identifications)
}

func (h *IdentificationController) GetAll(c echo.Context) (err error) {
	identifications, err := h.Repository.GetAll(
		c.Request().Context(),
		c.Request().Header.Get(enums.UserId),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for index := range identifications {
		identifications[index].Path = fmt.Sprintf(
			"%s/%s/%s",
			config.InitS3().Endpoint,
			config.InitS3().Buckets.EscortProfile,
			identifications[index].Path,
		)
	}

	return c.JSON(http.StatusOK, identifications)
}

func (h *IdentificationController) Create(c echo.Context) (err error) {
	image, _ := c.FormFile("image")

	if image.Size > constants.MaxImageSize {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	partId := c.FormValue("identificationPartId")
	ctx := c.Request().Context()

	if _, err := h.IdentificationCategoryRepository.GetById(ctx, partId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	src, _ := image.Open()
	defer src.Close()

	userId := c.Request().Header.Get(enums.UserId)

	url, err := h.S3Service.Upload(
		config.InitS3().Buckets.EscortProfile,
		image.Filename,
		userId,
		src,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	identification := models.Identification{
		Path:                 fmt.Sprintf("%s/%s", userId, image.Filename),
		ProfileId:            userId,
		IdentificationPartId: partId,
	}

	if err = identification.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(ctx, &identification); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	identification.Path = url

	return c.JSON(http.StatusCreated, identification)
}

func (h *IdentificationController) UpdateOne(c echo.Context) (err error) {
	id := c.Param("id")
	ctx := c.Request().Context()

	if _, err := h.IdentificationCategoryRepository.GetById(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	userId := c.Request().Header.Get(enums.UserId)
	identification, err := h.Repository.GetOne(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	image, _ := c.FormFile("image")

	if image.Size > constants.MaxImageSize {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	src, _ := image.Open()
	defer src.Close()

	url, err := h.S3Service.Upload(
		config.InitS3().Buckets.EscortProfile,
		image.Filename,
		userId,
		src,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	identification.Path = fmt.Sprintf("%s/%s", userId, image.Filename)

	if err = h.Repository.UpdateOne(ctx, userId, &identification); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	identification.Path = url

	return c.JSON(http.StatusOK, identification)
}
