package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapTimeCategoryRoutes(v *echo.Group) {
	router := &controllers.TimeCategoryController{
		Repository: &repositories.TimeCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/time-categories", router.GetAll)
}
