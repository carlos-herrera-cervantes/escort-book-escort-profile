package controllers

import (
	"encoding/json"
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
	var payload types.Payload

	json.NewDecoder(c.Request().Body).Decode(&payload)
	c.Bind(&pager)

	schedules, err := h.Repository.GetAll(c.Request().Context(), payload.User.Id, pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), payload.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, schedules))
}

func (h *ScheduleController) Create(c echo.Context) (err error) {
	var scheduleWrapper models.ScheduleWrapper

	c.Bind(&scheduleWrapper)
	schedule := scheduleWrapper.Map()

	if err = schedule.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), schedule); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, schedule)
}

func (h *ScheduleController) DeleteOne(c echo.Context) (err error) {
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
