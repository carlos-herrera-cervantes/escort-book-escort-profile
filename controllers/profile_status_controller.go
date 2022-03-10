package controllers

import (
	"encoding/json"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileStatusController struct {
	Repository *repositories.ProfileStatusRepository
}

func (h *ProfileStatusController) UpdateOne(c echo.Context) (err error) {
	var profileStatusWrapper models.ProfileStatusWrapper

	json.NewDecoder(c.Request().Body).Decode(&profileStatusWrapper)
	profileStatus, err := h.Repository.GetOne(c.Request().Context(), profileStatusWrapper.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profileStatus.ProfileStatusCategoryId = profileStatusWrapper.ProfileStatusCategoryId

	if err = h.Repository.UpdateOne(c.Request().Context(), profileStatusWrapper.User.Id, &profileStatus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profileStatus)
}
