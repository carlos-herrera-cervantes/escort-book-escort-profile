package controllers

import (
	"net/http"

	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type IdentificationPartController struct {
	Repository repositories.IIdentificationPartRepository
}

func (h *IdentificationPartController) GetAll(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	parts, err := h.Repository.GetAll(ctx, pager.Offset, pager.Limit)
	totalRows, _ := h.Repository.Count(ctx)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  parts,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}
