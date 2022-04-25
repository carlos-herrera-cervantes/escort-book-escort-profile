package controllers

import (
	"escort-book-escort-profile/constants"
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
	Repository                       repositories.IIdentificationRepository
	S3Service                        services.IS3Service
	IdentificationCategoryRepository repositories.IIdentificationPartRepository
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

	if image.Size > constants.MaxImageSize {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	partId := c.FormValue("identificationPartId")

	if _, err := h.IdentificationCategoryRepository.GetById(c.Request().Context(), partId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

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

	identification := models.Identification{
		Path:                 fmt.Sprintf("%s/%s", userId, image.Filename),
		ProfileId:            userId,
		IdentificationPartId: partId,
	}

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
	id := c.Param("id")

	if _, err := h.IdentificationCategoryRepository.GetById(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	userId := c.Request().Header.Get(enums.UserId)
	identification, err := h.Repository.GetOne(c.Request().Context(), userId)

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

	identification.Path = fmt.Sprintf("%s/%s", userId, image.Filename)

	if err = h.Repository.UpdateOne(c.Request().Context(), userId, &identification); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	identification.Path = url

	return c.JSON(http.StatusOK, identification)
}
