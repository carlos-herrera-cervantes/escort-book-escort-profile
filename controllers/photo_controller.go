package controllers

import (
	"escort-book-escort-profile/enums"
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
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId := c.Request().Header.Get(enums.UserId)
	photos, err := h.Repository.GetAll(c.Request().Context(), userId, pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, photos))
}

func (h *PhotoController) GetById(c echo.Context) (err error) {
	var pager types.Pager
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	photos, err := h.Repository.GetAll(c.Request().Context(), c.Param("id"), pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for index := range photos {
		photos[index].Path = fmt.Sprintf(
			"%s/%s/%s",
			os.Getenv("S3_ENPOINT"),
			os.Getenv("S3"), photos[index].Path,
		)
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, photos))
}

func (h *PhotoController) Create(c echo.Context) (err error) {
	image, _ := c.FormFile("image")
	src, _ := image.Open()

	defer src.Close()
	userId := c.Request().Header.Get(enums.UserId)

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

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
