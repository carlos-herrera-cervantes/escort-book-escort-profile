package controllers

import (
	"encoding/json"
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
	var payload types.Payload

	json.NewDecoder(c.Request().Body).Decode(&payload)
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	prices, err := h.Repository.GetAll(c.Request().Context(), payload.User.Id, pager.Offset, pager.Limit)
	number, _ := h.Repository.Count(c.Request().Context(), payload.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, prices))
}

func (h *PriceController) GetOne(c echo.Context) error {
	var payload types.Payload

	json.NewDecoder(c.Request().Body).Decode(&payload)
	price, err := h.Repository.GetOne(c.Request().Context(), payload.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, price)
}

func (h *PriceController) Create(c echo.Context) (err error) {
	var priceWrapper models.PriceWrapper

	c.Bind(&priceWrapper)
	price := models.Price{
		Cost:           priceWrapper.Cost,
		ProfileId:      priceWrapper.User.Id,
		TimeCategoryId: priceWrapper.TimeCategoryId,
		Quantity:       priceWrapper.Quantity,
	}

	if err = price.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, price)
}

func (h *PriceController) UpdateOne(c echo.Context) (err error) {
	var priceWrapper models.PriceWrapper

	c.Bind(&priceWrapper)
	price := models.Price{
		Cost:           priceWrapper.Cost,
		ProfileId:      priceWrapper.User.Id,
		TimeCategoryId: priceWrapper.TimeCategoryId,
		Quantity:       priceWrapper.Quantity,
	}

	if err = price.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, err = h.Repository.GetOne(c.Request().Context(), priceWrapper.User.Id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.UpdateOne(c.Request().Context(), priceWrapper.User.Id, &price); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, price)
}

func (h *PriceController) DeleteOne(c echo.Context) (err error) {
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
