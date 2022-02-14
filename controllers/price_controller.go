package controllers

import (
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PriceController struct {
	Repository *repositories.PriceRepository
}

func (h *PriceController) GetAll(c echo.Context) (err error) {
	var pager types.Pager

	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	prices, err := h.Repository.GetAll(c.Request().Context(), pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, prices))
}

func (h *PriceController) GetOne(c echo.Context) error {
	id := c.Param("id")
	price, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, price)
}

func (h *PriceController) Create(c echo.Context) (err error) {
	var price models.Price

	if err = c.Bind(&price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	id := c.Param("profileId")
	price.ProfileId = id

	if err = price.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, price)
}

func (h *PriceController) UpdateOne(c echo.Context) (err error) {
	var price models.Price
	id := c.Param("id")

	if err = c.Bind(&price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = price.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.UpdateOne(c.Request().Context(), id, &price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	price.Id = id

	return c.JSON(http.StatusOK, price)
}

func (h *PriceController) DeleteOne(c echo.Context) (err error) {
	id := c.Param("id")

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
