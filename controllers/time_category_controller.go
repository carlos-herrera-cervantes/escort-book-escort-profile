package controllers

import (
	"net/http"

	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type TimeCategoryController struct {
	Repository repositories.ITimeCategoryRepository
}

func (h *TimeCategoryController) GetAll(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	categories, err := h.Repository.GetAll(ctx, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalRows, _ := h.Repository.Count(ctx)
	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  categories,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}
