package controllers

import (
	"net/http"

	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

type BiographyController struct {
	Repository repositories.IBiographyRepository
}

func (h *BiographyController) GetOne(c echo.Context) error {
	biography, err := h.Repository.GetOne(
		c.Request().Context(),
		c.Request().Header.Get(enums.UserId),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, biography)
}

func (h *BiographyController) GetById(c echo.Context) error {
	biography, err := h.Repository.GetOne(c.Request().Context(), c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, biography)
}

func (h *BiographyController) Create(c echo.Context) (err error) {
	var biography models.Biography

	_ = c.Bind(&biography)
	biography.ProfileId = c.Request().Header.Get(enums.UserId)

	if err = biography.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &biography); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, biography)
}

func (h *BiographyController) UpdateOne(c echo.Context) (err error) {
	biography := models.Biography{}
	_ = c.Bind(&biography)

	ctx := c.Request().Context()
	userId := c.Request().Header.Get(enums.UserId)

	biography.ProfileId = userId

	if err = biography.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err = h.Repository.GetOne(ctx, userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.UpdateOne(ctx, userId, &biography); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, biography)
}

func (h *BiographyController) DeleteOne(c echo.Context) (err error) {
	userId := c.Request().Header.Get(enums.UserId)
	ctx := c.Request().Context()

	if _, err = h.Repository.GetOne(ctx, userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(ctx, userId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
