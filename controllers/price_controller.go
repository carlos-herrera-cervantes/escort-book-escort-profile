package controllers

import (
	"net/http"

	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type PriceController struct {
	Repository             repositories.IPriceRepository
	TimeCategoryRepository repositories.ITimeCategoryRepository
}

func (h *PriceController) GetAll(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	userId := c.Request().Header.Get(enums.UserId)
	prices, err := h.Repository.GetAll(ctx, userId, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalRows, _ := h.Repository.Count(ctx, userId)
	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  prices,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *PriceController) GetById(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	profileId := c.Param("id")
	prices, err := h.Repository.GetAll(ctx, profileId, pager.Offset, pager.Limit)
	totalRows, _ := h.Repository.Count(ctx, profileId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  prices,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *PriceController) Create(c echo.Context) (err error) {
	price := models.Price{}
	ctx := c.Request().Context()
	_ = c.Bind(&price)

	if _, err := h.TimeCategoryRepository.GetById(ctx, price.TimeCategoryId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	price.ProfileId = c.Request().Header.Get(enums.UserId)

	if err = price.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(ctx, &price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, price)
}

func (h *PriceController) UpdateOne(c echo.Context) (err error) {
	partialPrice := models.PricePartial{}
	ctx := c.Request().Context()
	_ = c.Bind(&partialPrice)

	if _, err := h.TimeCategoryRepository.GetById(ctx, partialPrice.TimeCategoryId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	userId := c.Request().Header.Get(enums.UserId)
	price, err := h.Repository.GetOne(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	partialPrice.MapPartial(&price)

	if err = h.Repository.UpdateOne(ctx, userId, &price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, price)
}

func (h *PriceController) DeleteOne(c echo.Context) (err error) {
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
