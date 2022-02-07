package controllers

import (
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
	id := c.Param("profileId")
	avatar, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, avatar)
}

func (h *AvatarController) Create(c echo.Context) (err error) {
	image, _ := c.FormFile("image")
	src, _ := image.Open()
	id := c.Param("profileId")

	defer src.Close()

	url, err := h.S3Service.Upload(c.Request().Context(), os.Getenv("S3"), image.Filename, id, src)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var avatar models.Avatar

	avatar.Path = fmt.Sprintf("%s/%s", id, image.Filename)
	avatar.ProfileId = id

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
	id := c.Param("profileId")
	avatar, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	image, _ := c.FormFile("image")
	src, _ := image.Open()

	defer src.Close()

	url, err := h.S3Service.Upload(c.Request().Context(), os.Getenv("S3"), image.Filename, id, src)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	avatar.Path = fmt.Sprintf("%s/%s", id, image.Filename)

	if err = h.Repository.UpdateOne(c.Request().Context(), id, &avatar); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	avatar.Path = url

	return c.JSON(http.StatusOK, avatar)
}

func (h *AvatarController) DeleteOne(c echo.Context) (err error) {
	id := c.Param("profileId")

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
