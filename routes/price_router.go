package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapPriceRoutes(v *echo.Group) {
	router := &controllers.PriceController{
		Repository: &repositories.PriceRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/profile/prices", router.GetAll)
	v.GET("/escort/profile/prices/:id", router.GetById)
	v.POST("/escort/profile/prices", router.Create)
	v.PUT("/escort/profile/prices/:id", router.UpdateOne)
	v.DELETE("/escort/profile/prices/:id", router.DeleteOne)
}
