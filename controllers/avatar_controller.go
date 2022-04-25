package controllers

import (
	"escort-book-escort-profile/constants"
	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type AvatarController struct {
	Repository *repositories.AvatarRepository
	S3Service  *services.S3Service
}

func (h *AvatarController) GetOne(c echo.Context) error {
	avatar, err := h.Repository.GetOne(c.Request().Context(), c.Request().Header.Get(enums.UserId))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	avatar.Path = fmt.Sprintf("%s/%s/%s", os.Getenv("S3_ENPOINT"), os.Getenv("S3"), avatar.Path)

	return c.JSON(http.StatusOK, avatar)
}

func (h *AvatarController) GetById(c echo.Context) (err error) {
	avatar, err := h.Repository.GetOne(c.Request().Context(), c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	avatar.Path = fmt.Sprintf("%s/%s/%s", os.Getenv("S3_ENPOINT"), os.Getenv("S3"), avatar.Path)

	return c.JSON(http.StatusOK, avatar)
}

func (h *AvatarController) Create(c echo.Context) (err error) {
	image, _ := c.FormFile("image")

	if image.Size > constants.MaxImageSize {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	src, _ := image.Open()
	userId := c.Request().Header.Get(enums.UserId)

	defer src.Close()

	url, err := h.S3Service.Upload(
		c.Request().Context(),
		os.Getenv("S3"),
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

	if err = h.Repository.Create(c.Request().Context(), &avatar); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	avatar.Path = url

	return c.JSON(http.StatusCreated, avatar)
}

func (h *AvatarController) UpdateOne(c echo.Context) (err error) {
	userId := c.Request().Header.Get(enums.UserId)
	avatar, err := h.Repository.GetOne(c.Request().Context(), userId)

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
		c.Request().Context(),
		os.Getenv("S3"),
		image.Filename,
		userId,
		src,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	avatar.Path = fmt.Sprintf("%s/%s", userId, image.Filename)

	if err = h.Repository.UpdateOne(c.Request().Context(), userId, &avatar); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	avatar.Path = url

	return c.JSON(http.StatusOK, avatar)
}

func (h *AvatarController) DeleteOne(c echo.Context) (err error) {
	userId := c.Request().Header.Get(enums.UserId)

	if _, err = h.Repository.GetOne(c.Request().Context(), userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), userId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
