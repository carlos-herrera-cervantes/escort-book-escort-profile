package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapServiceCategoryRoutes(v *echo.Group) {
	router := &controllers.ServiceCategoryController{
		Repository: &repositories.ServiceCategoryRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/service-categories", router.GetAll)
}
