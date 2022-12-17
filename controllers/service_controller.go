package controllers

import (
	"net/http"

	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type ServiceController struct {
	Repository                repositories.IServiceRepository
	ServiceCategoryRepository repositories.IServiceCategoryRepository
}

func (h *ServiceController) GetAll(c echo.Context) (err error) {
	pager := types.Pager{}
	_ = c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId := c.Request().Header.Get(enums.UserId)
	ctx := c.Request().Context()
	services, err := h.Repository.GetAll(ctx, userId, pager.Offset, pager.Limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	totalRows, _ := h.Repository.Count(ctx, userId)
	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  services,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *ServiceController) GetById(c echo.Context) (err error) {
	var pager types.Pager
	c.Bind(&pager)

	if err = pager.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId := c.Param("id")
	services, err := h.Repository.GetAll(c.Request().Context(), userId, pager.Offset, pager.Limit)
	totalRows, _ := h.Repository.Count(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagerResult := types.PagerResult{
		Pager: pager,
		Total: totalRows,
		Data:  services,
	}

	return c.JSON(http.StatusOK, pagerResult.Pages())
}

func (h *ServiceController) Create(c echo.Context) (err error) {
	service := models.Service{}
	ctx := c.Request().Context()
	_ = c.Bind(&service)

	if _, err := h.ServiceCategoryRepository.GetById(ctx, service.ServiceCategoryId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	service.ProfileId = c.Request().Header.Get(enums.UserId)

	if err = service.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.Repository.Create(ctx, &service); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, service)
}

func (h *ServiceController) DeleteOne(c echo.Context) (err error) {
	id := c.Param("id")
	userId := c.Request().Header.Get(enums.UserId)
	ctx := c.Request().Context()

	if _, err = h.Repository.GetOne(ctx, id, userId); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.Repository.DeleteOne(ctx, id, userId); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
