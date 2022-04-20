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

	v.GET("/escort/profile/biography", router.GetOne)
	v.GET("/escort/:id/profile/biography", router.GetById)
	v.POST("/escort/profile/biography", router.Create)
	v.PUT("/escort/profile/biography", router.UpdateOne)
	v.DELETE("/escort/profile/biography", router.DeleteOne)
}
