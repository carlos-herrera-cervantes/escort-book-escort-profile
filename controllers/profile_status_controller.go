package controllers

import (
	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileStatusController struct {
	Repository                      repositories.IProfileStatusRepository
	ProfileStatusCategoryRepository repositories.IProfileStatusCategoryRepository
}

func (h *ProfileStatusController) GetByExternal(c echo.Context) (err error) {
	id := c.Param("id")
	status, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, status)
}

func (h *ProfileStatusController) UpdateOne(c echo.Context) (err error) {
	var partialProfileStatus models.PartialProfileStatus
	c.Bind(&partialProfileStatus)

	category, err := h.ProfileStatusCategoryRepository.GetById(c.Request().Context(), partialProfileStatus.ProfileStatusCategoryId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	validStatus := map[string]bool{enums.Deactivated: true, enums.Deleted: true}
	_, ok := validStatus[category.Name]

	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	userId := c.Request().Header.Get(enums.UserId)
	profileStatus, err := h.Repository.GetOne(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profileStatus.ProfileStatusCategoryId = partialProfileStatus.ProfileStatusCategoryId

	if err = h.Repository.UpdateOne(c.Request().Context(), userId, &profileStatus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profileStatus)
}
