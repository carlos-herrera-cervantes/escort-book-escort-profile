package controllers

import (
	"escort-book-escort-profile/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	Repository *repositories.ProfileRepository
}

func (h *ProfileController) GetOne(c echo.Context) error {
	id := c.Param("id")
	profile, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return c.NoContent(http.StatusOK)
	}

	return c.JSON(http.StatusOK, profile)
}
