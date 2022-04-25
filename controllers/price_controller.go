package controllers

import (
	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PriceController struct {
	Repository             repositories.IPriceRepository
	TimeCategoryRepository repositories.ITimeCategoryRepository
}

func (h *PriceController) GetAll(c echo.Context) (err error) {
	var pager types.Pager
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId := c.Request().Header.Get(enums.UserId)
	prices, err := h.Repository.GetAll(c.Request().Context(), userId, pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, prices))
}

func (h *PriceController) GetById(c echo.Context) (err error) {
	var pager types.Pager
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	prices, err := h.Repository.GetAll(c.Request().Context(), c.Param("id"), pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, prices))
}

func (h *PriceController) Create(c echo.Context) (err error) {
	var price models.Price
	c.Bind(&price)

	if _, err := h.TimeCategoryRepository.GetById(c.Request().Context(), price.TimeCategoryId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	price.ProfileId = c.Request().Header.Get(enums.UserId)

	if err = price.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, price)
}

func (h *PriceController) UpdateOne(c echo.Context) (err error) {
	var partialPrice models.PricePartial
	c.Bind(&partialPrice)

	if _, err := h.TimeCategoryRepository.GetById(c.Request().Context(), partialPrice.TimeCategoryId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	userId := c.Request().Header.Get(enums.UserId)
	price, err := h.Repository.GetOne(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	partialPrice.MapPartial(&price)

	if err = h.Repository.UpdateOne(c.Request().Context(), userId, &price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, price)
}

func (h *PriceController) DeleteOne(c echo.Context) (err error) {
	userId := c.Request().Header.Get(enums.UserId)

	if _, err = h.Repository.GetOne(c.Request().Context(), userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), userId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
