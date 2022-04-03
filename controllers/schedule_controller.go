package controllers

import (
	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ScheduleController struct {
	Repository *repositories.ScheduleRepository
}

func (h *ScheduleController) GetAll(c echo.Context) (err error) {
	var pager types.Pager
	c.Bind(&pager)

	userId := c.Request().Header.Get(enums.UserId)
	schedules, err := h.Repository.GetAll(c.Request().Context(), userId, pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, schedules))
}

func (h *ScheduleController) GetById(c echo.Context) (err error) {
	var pager types.Pager
	c.Bind(&pager)

	schedules, err := h.Repository.GetAll(c.Request().Context(), c.Param("id"), pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, schedules))
}

func (h *ScheduleController) Create(c echo.Context) (err error) {
	var schedule models.Schedule

	c.Bind(&schedule)

	if err = schedule.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &schedule); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, schedule)
}

func (h *ScheduleController) DeleteOne(c echo.Context) (err error) {
	userId := c.Request().Header.Get(enums.UserId)

	if _, err = h.Repository.GetOne(c.Request().Context(), userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), userId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
