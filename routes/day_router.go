package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapDayRoutes(v *echo.Group) {
	router := &controllers.DayController{
		Repository: &repositories.DayRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/days", router.GetAll)
}
