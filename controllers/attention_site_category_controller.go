package controllers

import (
	"net/http"

	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type AttentionSiteCategoryController struct {
	Repository repositories.IAttentionSiteCategoryRepository
}

func (h *AttentionSiteCategoryController) GetAll(c echo.Context) (err error) {
	var pager types.Pager

	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	categories, err := h.Repository.GetAll(c.Request().Context(), pager.Offset, pager.Limit)
	totalRows, _ := h.Repository.Count(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  categories,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}
