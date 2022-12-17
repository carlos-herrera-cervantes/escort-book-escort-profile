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
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type PhotoController struct {
	Repository repositories.IPhotoRepository
	S3Service  services.IS3Service
}

func (h *PhotoController) GetAll(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	userId := c.Request().Header.Get(enums.UserId)
	photos, err := h.Repository.GetAll(ctx, userId, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for index := range photos {
		photos[index].Path = fmt.Sprintf(
			"%s/%s/%s",
			config.InitS3().Endpoint,
			config.InitS3().Buckets.EscortProfile,
			photos[index].Path,
		)
	}

	totalRows, _ := h.Repository.Count(ctx, userId)
	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  photos,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *PhotoController) GetById(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	profileId := c.Param("id")
	photos, err := h.Repository.GetAll(ctx, profileId, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalRows, _ := h.Repository.Count(ctx, profileId)

	for index := range photos {
		photos[index].Path = fmt.Sprintf(
			"%s/%s/%s",
			config.InitS3().Endpoint,
			config.InitS3().Buckets.EscortProfile,
			photos[index].Path,
		)
	}

	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  photos,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *PhotoController) Create(c echo.Context) (err error) {
	image, _ := c.FormFile("image")

	if image.Size > constants.MaxImageSize {
		return echo.NewHTTPError(http.StatusBadRequest)
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

	photo := models.Photo{
		Path:      fmt.Sprintf("%s/%s", userId, image.Filename),
		ProfileId: userId,
	}

	if err = photo.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &photo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	photo.Path = url

	return c.JSON(http.StatusCreated, photo)
}

func (h *PhotoController) DeleteOne(c echo.Context) (err error) {
	id := c.Param("id")
	ctx := c.Request().Context()

	if _, err = h.Repository.GetOne(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
