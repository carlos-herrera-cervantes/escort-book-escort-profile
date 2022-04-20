package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapServiceRoutes(v *echo.Group) {
	router := &controllers.ServiceController{
		Repository: &repositories.ServiceRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/profile/service", router.GetAll)
	v.GET("/escort/:id/profile/service", router.GetById)
	v.DELETE("/escort/profile/service/:id", router.DeleteOne)
	v.POST("/escort/profile/service", router.Create)
}
