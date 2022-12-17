package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapBiographyRoutes(v *echo.Group) {
	router := &controllers.BiographyController{
		Repository: &repositories.BiographyRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/profile/biography", router.GetOne)
	v.GET("/escort/:id/profile/biography", router.GetById)
	v.POST("/escort/profile/biography", router.Create)
	v.PUT("/escort/profile/biography", router.UpdateOne)
	v.DELETE("/escort/profile/biography", router.DeleteOne)
}
