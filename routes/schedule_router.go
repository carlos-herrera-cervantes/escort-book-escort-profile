package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapScheduleRoutes(v *echo.Group) {
	router := &controllers.ScheduleController{
		Repository: &repositories.ScheduleRepository{
			Data: singleton.NewPostgresClient(),
		},
		DayRepository: &repositories.DayRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/profile/schedules", router.GetAll)
	v.GET("/escort/:id/profile/schedules", router.GetById)
	v.POST("/escort/profile/schedules", router.Create)
	v.DELETE("/escort/profile/schedules/:id", router.DeleteOne)
}
