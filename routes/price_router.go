package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapPriceRoutes(v *echo.Group) {
	router := &controllers.PriceController{
		Repository: &repositories.PriceRepository{
			Data: singleton.NewPostgresClient(),
		},
		TimeCategoryRepository: &repositories.TimeCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/profile/prices", router.GetAll)
	v.GET("/escort/:id/profile/prices", router.GetById)
	v.POST("/escort/profile/prices", router.Create)
	v.PUT("/escort/profile/prices/:id", router.UpdateOne)
	v.DELETE("/escort/profile/prices/:id", router.DeleteOne)
}
