package controllers

import (
	"encoding/json"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BiographyController struct {
	Repository *repositories.BiographyRepository
}

func (h *BiographyController) GetOne(c echo.Context) error {
	var payload types.Payload

	json.NewDecoder(c.Request().Body).Decode(&payload)
	biography, err := h.Repository.GetOne(c.Request().Context(), payload.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, biography)
}

func (h *BiographyController) Create(c echo.Context) (err error) {
	var biographyWrapper models.BiographyWrapper

	if err = c.Bind(&biographyWrapper); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	biography := models.Biography{
		ProfileId:   biographyWrapper.User.Id,
		Description: biographyWrapper.Description,
	}

	if err = biography.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &biography); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, biography)
}

func (h *BiographyController) UpdateOne(c echo.Context) (err error) {
	var biographyWrapper models.BiographyWrapper

	if err = c.Bind(&biographyWrapper); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	biography := models.Biography{
		ProfileId:   biographyWrapper.User.Id,
		Description: biographyWrapper.Description,
	}

	if err = biography.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err = h.Repository.GetOne(c.Request().Context(), biographyWrapper.User.Id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.UpdateOne(c.Request().Context(), biographyWrapper.User.Id, &biography); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, biography)
}

func (h *BiographyController) DeleteOne(c echo.Context) (err error) {
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
