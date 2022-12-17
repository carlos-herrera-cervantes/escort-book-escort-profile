package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapServiceRoutes(v *echo.Group) {
	router := &controllers.ServiceController{
		Repository: &repositories.ServiceRepository{
			Data: singleton.NewPostgresClient(),
		},
		ServiceCategoryRepository: &repositories.ServiceCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/profile/service", router.GetAll)
	v.GET("/escort/:id/profile/service", router.GetById)
	v.DELETE("/escort/profile/service/:id", router.DeleteOne)
	v.POST("/escort/profile/service", router.Create)
}
