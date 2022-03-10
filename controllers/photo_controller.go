package controllers

import (
	"encoding/json"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/types"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type PhotoController struct {
	Repository *repositories.PhotoRepository
	S3Service  *services.S3Service
}

func (h *PhotoController) GetAll(c echo.Context) (err error) {
	var pager types.Pager
	var payload types.Payload

	json.NewDecoder(c.Request().Body).Decode(&payload)
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	photos, err := h.Repository.GetAll(c.Request().Context(), payload.User.Id, pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), payload.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, photos))
}

func (h *PhotoController) Create(c echo.Context) (err error) {
	image, _ := c.FormFile("image")
	src, _ := image.Open()

	defer src.Close()

	var photoWrapper models.PhotoWrapper
	json.NewDecoder(c.Request().Body).Decode(&photoWrapper)

	url, err := h.S3Service.Upload(
		c.Request().Context(),
		os.Getenv("S3"),
		image.Filename,
		photoWrapper.User.Id,
		src,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	photo := models.Photo{
		Path:      fmt.Sprintf("%s/%s", photoWrapper.User.Id, image.Filename),
		ProfileId: photoWrapper.User.Id,
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

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
