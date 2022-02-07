package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapTimeCategoryRoutes(v *echo.Group) {
	router := &controllers.TimeCategoryController{
		Repository: &repositories.TimeCategoryRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/time-categories", router.GetAll)
}
