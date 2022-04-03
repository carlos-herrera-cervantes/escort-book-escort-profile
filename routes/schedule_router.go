package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapScheduleRoutes(v *echo.Group) {
	router := &controllers.ScheduleController{
		Repository: &repositories.ScheduleRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/profile/schedules", router.GetAll)
	v.GET("/escort/profile/schedules/:id", router.GetById)
	v.POST("/escort/profile/schedules", router.Create)
	v.DELETE("/escort/profile/schedules/:id", router.DeleteOne)
}
