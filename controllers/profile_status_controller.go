package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

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
	partialProfileStatus := models.PartialProfileStatus{}
	ctx := c.Request().Context()
	_ = c.Bind(&partialProfileStatus)

	category, err := h.ProfileStatusCategoryRepository.GetById(
		ctx,
		partialProfileStatus.ProfileStatusCategoryId,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}

	userId := c.Param("id")
	profileStatus, err := h.Repository.GetOne(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profileStatus.ProfileStatusCategoryId = partialProfileStatus.ProfileStatusCategoryId

	if err = h.Repository.UpdateOne(ctx, userId, &profileStatus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	args := map[string]string{
		"userId":       userId,
		"userType":     c.Request().Header.Get("user-type"),
		"userEmail":    c.Request().Header.Get("user-email"),
		"categoryName": category.Name,
	}
	go h.emitDeletionMessage(ctx, &wg, args)
	go h.emitDisableMessage(ctx, &wg, args)

	wg.Wait()

	return c.JSON(http.StatusOK, profileStatus)
}

func (h *ProfileStatusController) UpdateOne(c echo.Context) (err error) {
	partialProfileStatus := models.PartialProfileStatus{}
	ctx := c.Request().Context()
	_ = c.Bind(&partialProfileStatus)

	category, err := h.ProfileStatusCategoryRepository.GetById(
		ctx,
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
	profileStatus, err := h.Repository.GetOne(ctx, userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	profileStatus.ProfileStatusCategoryId = partialProfileStatus.ProfileStatusCategoryId

	if err = h.Repository.UpdateOne(ctx, userId, &profileStatus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	args := map[string]string{
		"userId":       userId,
		"userType":     c.Request().Header.Get("user-type"),
		"userEmail":    c.Request().Header.Get("user-email"),
		"categoryName": category.Name,
	}
	go h.emitDisableMessage(ctx, &wg, args)
	go h.emitDeletionMessage(ctx, &wg, args)

	wg.Wait()

	return c.JSON(http.StatusOK, profileStatus)
}

func (h *ProfileStatusController) emitDisableMessage(
	ctx context.Context,
	wg *sync.WaitGroup,
	args map[string]string,
) {
	defer wg.Done()

	if args["categoryName"] != enums.Deleted && args["categoryName"] != enums.Deactivated {
		return
	}

	blockUserEvent := types.BlockUserEvent{
		UserId: args["userId"],
		Status: args["categoryName"],
	}
	message, _ := json.Marshal(blockUserEvent)

	if err := h.KafkaService.SendMessage("block-user", message); err != nil {
		log.Println("WE CAN'T PROPAGATE THE PROFILE STATUS: ", err.Error())
	}
}

func (h *ProfileStatusController) emitDeletionMessage(
	ctx context.Context,
	wg *sync.WaitGroup,
	args map[string]string,
) {
	defer wg.Done()

	if args["categoryName"] != enums.Deleted {
		return
	}

	deleteUserEvent := types.DeleteUserEvent{
		UserId:    args["userId"],
		UserType:  args["userType"],
		UserEmail: args["userEmail"],
	}
	message, _ := json.Marshal(deleteUserEvent)

	if err := h.KafkaService.SendMessage("user-delete-account", message); err != nil {
		log.Println("WE CAN'T PROPAGATE DELETION OF CUSTOMER: ", err.Error())
	}
}
