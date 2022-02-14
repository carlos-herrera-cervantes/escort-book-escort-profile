package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapDayRoutes(v *echo.Group) {
	router := &controllers.DayController{
		Repository: &repositories.DayRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/days", router.GetAll)
}
