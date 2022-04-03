package controllers

import (
	"escort-book-escort-profile/enums"
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
	profile, err := h.Repository.GetOne(c.Request().Context(), c.Request().Header.Get(enums.UserId))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileController) GetById(c echo.Context) error {
	profile, err := h.Repository.GetOne(c.Request().Context(), c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileController) Create(c echo.Context) (err error) {
	var profile models.Profile
	c.Bind(&profile)

	profile.EscortId = c.Request().Header.Get(enums.UserId)

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
	var profilePartial models.PartialProfile
	c.Bind(&profilePartial)

	userId := c.Request().Header.Get(enums.UserId)
	profile, err := h.Repository.GetOne(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profilePartial.MapPartial(&profile)

	if err = h.Repository.UpdateOne(c.Request().Context(), userId, &profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileController) DeleteOne(c echo.Context) (err error) {
	userId := c.Request().Header.Get(enums.UserId)

	if _, err = h.Repository.GetOne(c.Request().Context(), userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), userId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
