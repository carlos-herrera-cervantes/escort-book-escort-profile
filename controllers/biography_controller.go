package controllers

import (
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BiographyController struct {
	Repository *repositories.BiographyRepository
}

func (h *BiographyController) GetOne(c echo.Context) error {
	id := c.Param("profileId")
	biography, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, biography)
}

func (h *BiographyController) Create(c echo.Context) (err error) {
	var biography models.Biography

	if err = c.Bind(&biography); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	profileId := c.Param("profileId")
	biography.ProfileId = profileId

	if err = biography.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &biography); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, biography)
}

func (h *BiographyController) UpdateOne(c echo.Context) (err error) {
	var biography models.Biography
	id := c.Param("profileId")

	if err = c.Bind(&biography); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = biography.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.UpdateOne(c.Request().Context(), id, &biography); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	biography.Id = id

	return c.JSON(http.StatusOK, biography)
}

func (h *BiographyController) DeleteOne(c echo.Context) (err error) {
	id := c.Param("profileId")

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
