package controllers

import (
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	Repository *repositories.ProfileRepository
	Emitter    *services.EmitterService
}

func (h *ProfileController) GetOne(c echo.Context) error {
	id := c.Param("id")
	profile, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileController) Create(c echo.Context) (err error) {
	var profile models.Profile

	if err = c.Bind(&profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = profile.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Emitter.Emit("create.profile.status", profile)

	return c.JSON(http.StatusCreated, profile)
}

func (h *ProfileController) UpdateOne(c echo.Context) (err error) {
	var profile models.Profile
	id := c.Param("id")

	if err = c.Bind(&profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = profile.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.UpdateOne(c.Request().Context(), id, &profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	profile.Id = id

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileController) DeleteOne(c echo.Context) (err error) {
	id := c.Param("id")

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
