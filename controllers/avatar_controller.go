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

type AvatarController struct {
	Repository repositories.IAvatarRepository
	S3Service  services.IS3Service
}

func (h *AvatarController) GetOne(c echo.Context) error {
	avatar, err := h.Repository.GetOne(c.Request().Context(), c.Request().Header.Get(enums.UserId))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	avatar.Path = fmt.Sprintf(
		"%s/%s/%s",
		config.InitS3().Endpoint,
		config.InitS3().Buckets.EscortProfile,
		avatar.Path,
	)

	return c.JSON(http.StatusOK, avatar)
}

func (h *AvatarController) GetById(c echo.Context) (err error) {
	avatar, err := h.Repository.GetOne(c.Request().Context(), c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	avatar.Path = fmt.Sprintf(
		"%s/%s/%s",
		config.InitS3().Endpoint,
		config.InitS3().Buckets.EscortProfile,
		avatar.Path,
	)

	return c.JSON(http.StatusOK, avatar)
}

func (h *AvatarController) Upsert(c echo.Context) (err error) {
	image, _ := c.FormFile("image")

	if image.Size > constants.MaxImageSize {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	src, _ := image.Open()
	userId := c.Request().Header.Get(enums.UserId)

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

	var avatar models.Avatar

	avatar.Path = fmt.Sprintf("%s/%s", userId, image.Filename)
	avatar.ProfileId = userId

	if err = avatar.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	if _, err = h.Repository.GetOne(ctx, userId); err != nil {
		if err = h.Repository.Create(ctx, &avatar); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	} else {
		if err = h.Repository.UpdateOne(ctx, userId, &avatar); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	avatar.Path = url

	return c.JSON(http.StatusCreated, avatar)
}

func (h *AvatarController) DeleteOne(c echo.Context) (err error) {
	ctx := c.Request().Context()
	userId := c.Request().Header.Get(enums.UserId)

	if _, err = h.Repository.GetOne(ctx, userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(ctx, userId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
