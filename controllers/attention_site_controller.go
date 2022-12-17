package controllers

import (
	"net/http"

	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type AttentionSiteController struct {
	Repository                      repositories.IAttentionSiteRepository
	AttentionSiteCategoryRepository repositories.IAttentionSiteCategoryRepository
}

func (h *AttentionSiteController) GetAll(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var pager types.Pager
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId := c.Request().Header.Get(enums.UserId)

	sites, err := h.Repository.GetAll(ctx, userId, pager.Offset, pager.Limit)
	totalRows, _ := h.Repository.Count(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  sites,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *AttentionSiteController) GetById(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")

	var pager types.Pager
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sites, err := h.Repository.GetAll(ctx, id, pager.Offset, pager.Limit)
	totalRows, _ := h.Repository.Count(ctx, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  sites,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *AttentionSiteController) Create(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var site models.AttentionSite
	_ = c.Bind(&site)

	if _, err := h.AttentionSiteCategoryRepository.GetById(ctx, site.AttentionSiteCategoryId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	site.ProfileId = c.Request().Header.Get(enums.UserId)

	if err = site.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(ctx, &site); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, site)
}

func (h *AttentionSiteController) DeleteOne(c echo.Context) (err error) {
	id := c.Param("id")
	ctx := c.Request().Context()

	if _, err = h.Repository.GetOne(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
