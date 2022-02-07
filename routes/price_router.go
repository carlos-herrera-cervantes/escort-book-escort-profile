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

	v.GET("/escort/profile/:profileId/prices/:id", router.GetOne)
	v.GET("/escort/profile/:profileId/prices", router.GetAll)
	v.POST("/escort/profile/:profileId/prices", router.Create)
	v.PUT("/escort/profile/:profileId/prices/:id", router.UpdateOne)
	v.DELETE("/escort/profile/:profileId/prices/:id", router.DeleteOne)
}
