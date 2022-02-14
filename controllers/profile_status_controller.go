package controllers

import (
	"escort-book-escort-profile/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileStatusController struct {
	Repository *repositories.ProfileStatusRepository
}

func (h *ProfileStatusController) UpdateOne(c echo.Context) (err error) {
	id := c.Param("profileId")

	profileStatus, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = c.Bind(&profileStatus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = h.Repository.UpdateOne(c.Request().Context(), id, &profileStatus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profileStatus)
}
