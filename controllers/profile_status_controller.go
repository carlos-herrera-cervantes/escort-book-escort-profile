package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"escort-book-escort-profile/enums"
	"escort-book-escort-profile/models"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/types"

	"github.com/labstack/echo/v4"
)

type ProfileStatusController struct {
	Repository                      repositories.IProfileStatusRepository
	ProfileStatusCategoryRepository repositories.IProfileStatusCategoryRepository
	KafkaService                    services.IKafkaService
}

func (h *ProfileStatusController) GetByExternal(c echo.Context) (err error) {
	id := c.Param("id")
	status, err := h.Repository.GetOne(c.Request().Context(), id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, status)
}

func (h *ProfileStatusController) UpdateByExternal(c echo.Context) (err error) {
	var partialProfileStatus models.PartialProfileStatus
	_ = c.Bind(&partialProfileStatus)

	category, err := h.ProfileStatusCategoryRepository.GetById(
		c.Request().Context(),
		partialProfileStatus.ProfileStatusCategoryId,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}

	userId := c.Param("id")
	profileStatus, err := h.Repository.GetOne(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profileStatus.ProfileStatusCategoryId = partialProfileStatus.ProfileStatusCategoryId

	if err = h.Repository.UpdateOne(c.Request().Context(), userId, &profileStatus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if category.Name == enums.Locked || category.Name == enums.Active {
		go h.emitDisableMessage(c.Request().Context(), userId, category.Name)
	}

	if category.Name == enums.Deleted {
		args := map[string]string{
			"userId": userId,
			"userType": c.Request().Header.Get("user-type"),
			"userEmail": c.Request().Header.Get("user-email"),
		}
		go h.emitDeletionMessage(c.Request().Context(), args)
		go h.emitDisableMessage(c.Request().Context(), userId, category.Name)
	}

	return c.JSON(http.StatusOK, profileStatus)
}

func (h *ProfileStatusController) UpdateOne(c echo.Context) (err error) {
	var partialProfileStatus models.PartialProfileStatus
	_ = c.Bind(&partialProfileStatus)

	category, err := h.ProfileStatusCategoryRepository.GetById(
		c.Request().Context(),
		partialProfileStatus.ProfileStatusCategoryId,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}

	validStatus := map[string]bool{enums.Deactivated: true, enums.Deleted: true}
	_, ok := validStatus[category.Name]

	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid operation")
	}

	userId := c.Request().Header.Get(enums.UserId)
	profileStatus, err := h.Repository.GetOne(c.Request().Context(), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profileStatus.ProfileStatusCategoryId = partialProfileStatus.ProfileStatusCategoryId

	if err = h.Repository.UpdateOne(c.Request().Context(), userId, &profileStatus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if category.Name == enums.Deactivated {
		go h.emitDisableMessage(c.Request().Context(), userId, category.Name)
	}

	if category.Name == enums.Deleted {
		args := map[string]string{
			"userId": userId,
			"userType": c.Request().Header.Get("user-type"),
			"userEmail": c.Request().Header.Get("user-email"),
		}
		go h.emitDeletionMessage(c.Request().Context(), args)
		go h.emitDisableMessage(c.Request().Context(), userId, category.Name)
	}

	return c.JSON(http.StatusOK, profileStatus)
}

func (h *ProfileStatusController) emitDisableMessage(ctx context.Context, userId, categoryName string) {
	newBlockUserEvent := types.BlockUserEvent{
		UserId: userId,
		Status: categoryName,
	}
	message, _ := json.Marshal(newBlockUserEvent)

	if err := h.KafkaService.SendMessage(ctx, "block-user", message); err != nil {
		log.Println("WE CAN'T PROPAGATE THE PROFILE STATUS: ", err.Error())
	}
}

func (h *ProfileStatusController) emitDeletionMessage(ctx context.Context, args map[string]string) {
	deleteUserEvent := types.DeleteUserEvent{
		UserId: args["userId"],
		UserType: args["userType"],
		UserEmail: args["userEmail"],
	}
	message, _ := json.Marshal(deleteUserEvent)

	if err := h.KafkaService.SendMessage(ctx, "user-delete-account", message); err != nil {
		log.Println("WE CAN'T PROPAGATE DELETION OF CUSTOMER: ", err.Error())
	}
}
