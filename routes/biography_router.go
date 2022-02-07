package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapBiographyRoutes(v *echo.Group) {
	router := &controllers.BiographyController{
		Repository: &repositories.BiographyRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/profile/:profileId/biography", router.GetOne)
	v.POST("/escort/profile/:profileId/biography", router.Create)
	v.PUT("/escort/profile/:profileId/biography", router.UpdateOne)
	v.DELETE("/escort/profile/:profileId/biography", router.DeleteOne)
}
