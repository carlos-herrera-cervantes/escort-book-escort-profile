package controllers

import (
	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	Repository            repositories.IProfileRepository
	Emitter               services.IEmitterService
	NationalityRepository repositories.INationalityRepository
}

func (h *ProfileController) GetAll(c echo.Context) (err error) {
	var pager types.Pager
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	escorts, err := h.Repository.GetAll(c.Request().Context(), pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	number, _ := h.Repository.Count(c.Request().Context())
	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, escorts))
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

	if _, err := h.NationalityRepository.GetById(c.Request().Context(), profile.NationalityId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

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

	if _, err := h.NationalityRepository.GetById(c.Request().Context(), profilePartial.NationalityId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

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
