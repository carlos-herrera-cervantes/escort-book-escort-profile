package controllers

import (
	"encoding/json"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	Repository *repositories.ProfileRepository
	Emitter    *services.EmitterService
}

func (h *ProfileController) GetOne(c echo.Context) error {
	var payload types.Payload

	json.NewDecoder(c.Request().Body).Decode(&payload)
	profile, err := h.Repository.GetOne(c.Request().Context(), payload.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileController) Create(c echo.Context) (err error) {
	var profileWrapper models.ProfileWrapper

	c.Bind(&profileWrapper)
	profile := profileWrapper.Map()

	if err = profile.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Emitter.Emit("create.profile.status", profileWrapper)

	return c.JSON(http.StatusCreated, profileWrapper)
}

func (h *ProfileController) UpdateOne(c echo.Context) (err error) {
	var profileWrapper models.ProfilePartialWrapper

	json.NewDecoder(c.Request().Body).Decode(&profileWrapper)
	finded, err := h.Repository.GetOne(c.Request().Context(), profileWrapper.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profile := profileWrapper.MapPartial(&finded)

	if err = h.Repository.UpdateOne(c.Request().Context(), profileWrapper.User.Id, profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileController) DeleteOne(c echo.Context) (err error) {
	var payload types.Payload

	json.NewDecoder(c.Request().Body).Decode(&payload)

	if _, err = h.Repository.GetOne(c.Request().Context(), payload.User.Id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), payload.User.Id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
