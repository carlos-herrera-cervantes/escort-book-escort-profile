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

type IdentificationController struct {
	Repository *repositories.IdentificationRepository
	S3Service  *services.S3Service
}

func (h *IdentificationController) GetAll(c echo.Context) (err error) {
	var pager types.Pager
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	identifications, err := h.Repository.GetAll(
		c.Request().Context(),
		c.Request().Header.Get(enums.UserId),
		pager.Offset,
		pager.Limit,
	)
	number, _ := h.Repository.Count(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, identifications))
}

func (h *IdentificationController) Create(c echo.Context) (err error) {
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

	var identification models.Identification
	partId := c.FormValue("identificationPartId")

	identification.Path = fmt.Sprintf("%s/%s", userId, image.Filename)
	identification.ProfileId = userId
	identification.IdentificationPartId = partId

	if err = identification.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &identification); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	identification.Path = url

	return c.JSON(http.StatusCreated, identification)
}

func (h *IdentificationController) UpdateOne(c echo.Context) (err error) {
	userId := c.Request().Header.Get(enums.UserId)
	identification, err := h.Repository.GetOne(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	image, _ := c.FormFile("image")
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

	identification.Path = fmt.Sprintf("%s/%s", userId, image.Filename)

	if err = h.Repository.UpdateOne(c.Request().Context(), userId, &identification); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	identification.Path = url

	return c.JSON(http.StatusOK, identification)
}
