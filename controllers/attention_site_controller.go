package controllers

import (
	"encoding/json"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AttentionSiteController struct {
	Repository *repositories.AttentionSiteRepository
}

func (h *AttentionSiteController) GetAll(c echo.Context) (err error) {
	var pager types.Pager
	var payload types.Payload

	json.NewDecoder(c.Request().Body).Decode(&payload)
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sites, err := h.Repository.GetAll(
		c.Request().Context(),
		payload.User.Id,
		pager.Offset,
		pager.Limit,
	)
	number, _ := h.Repository.Count(c.Request().Context(), payload.User.Id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{}

	return c.JSON(http.StatusOK, pagerResult.GetPagerResult(&pager, number, sites))
}

func (h *AttentionSiteController) Create(c echo.Context) (err error) {
	var wrapper models.AttentionSiteWrapper

	json.NewDecoder(c.Request().Body).Decode(&wrapper)
	site := models.AttentionSite{
		ProfileId:               wrapper.User.Id,
		AttentionSiteCategoryId: wrapper.AttentionSiteCategoryId,
	}

	if err = site.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(c.Request().Context(), &site); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, site)
}

func (h *AttentionSiteController) DeleteOne(c echo.Context) (err error) {
	id := c.Param("id")

	if _, err = h.Repository.GetOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
