package controllers

import (
	"net/http"

	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	Repository            repositories.IProfileRepository
	Emitter               services.IEmitterService
	NationalityRepository repositories.INationalityRepository
}

func (h *ProfileController) GetAll(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	escorts, err := h.Repository.GetAll(ctx, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalRows, _ := h.Repository.Count(ctx)
	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  escorts,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *ProfileController) GetOne(c echo.Context) error {
	profile, err := h.Repository.GetOne(
		c.Request().Context(),
		c.Request().Header.Get(enums.UserId),
	)

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
	profile := models.Profile{}
	ctx := c.Request().Context()
	_ = c.Bind(&profile)

	if _, err := h.NationalityRepository.GetById(ctx, profile.NationalityId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	profile.EscortId = c.Request().Header.Get(enums.UserId)

	if err = profile.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(ctx, &profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Emitter.Emit("create.profile.status", profile)

	return c.JSON(http.StatusCreated, profile)
}

func (h *ProfileController) UpdateOne(c echo.Context) (err error) {
	profilePartial := models.PartialProfile{}
	_ = c.Bind(&profilePartial)

	userId := c.Request().Header.Get(enums.UserId)
	ctx := c.Request().Context()
	profile, err := h.Repository.GetOne(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profilePartial.MapPartial(&profile)

	if err = h.Repository.UpdateOne(ctx, userId, &profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileController) DeleteOne(c echo.Context) (err error) {
	userId := c.Request().Header.Get(enums.UserId)
	ctx := c.Request().Context()

	if _, err = h.Repository.GetOne(ctx, userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(ctx, userId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
