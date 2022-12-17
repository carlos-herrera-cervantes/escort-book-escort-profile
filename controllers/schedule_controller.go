package controllers

import (
	"net/http"

	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type ScheduleController struct {
	Repository    repositories.IScheduleRepository
	DayRepository repositories.IDayRepository
}

func (h *ScheduleController) GetAll(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err := pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	userId := c.Request().Header.Get(enums.UserId)
	schedules, err := h.Repository.GetAll(ctx, userId, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalRows, _ := h.Repository.Count(ctx, userId)
	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  schedules,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *ScheduleController) GetById(c echo.Context) (err error) {
	pager := types.Pager{}
	ctx := c.Request().Context()
	_ = c.Bind(&pager)

	profileId := c.Param("id")
	schedules, err := h.Repository.GetAll(ctx, profileId, pager.Offset, pager.Limit)
	totalRows, _ := h.Repository.Count(ctx, profileId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  schedules,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *ScheduleController) Create(c echo.Context) (err error) {
	schedule := models.Schedule{}
	ctx := c.Request().Context()
	_ = c.Bind(&schedule)

	if _, err := h.DayRepository.GetById(ctx, schedule.DayId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	schedule.ProfileId = c.Request().Header.Get(enums.UserId)

	if err = schedule.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(ctx, &schedule); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, schedule)
}

func (h *ScheduleController) DeleteOne(c echo.Context) (err error) {
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
