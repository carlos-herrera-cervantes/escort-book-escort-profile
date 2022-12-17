package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapServiceCategoryRoutes(v *echo.Group) {
	router := &controllers.ServiceCategoryController{
		Repository: &repositories.ServiceCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/service-categories", router.GetAll)
}
