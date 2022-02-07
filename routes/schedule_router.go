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

	v.GET("/escort/profile/:profileId/schedules", router.GetAll)
	v.POST("/escort/profile/:profileId/schedules", router.Create)
	v.DELETE("/escort/profile/:profileId/schedules/:id", router.DeleteOne)
}
