package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapIdentificationPartRoutes(v *echo.Group) {
	router := &controllers.IdentificationPartController{
		Repository: &repositories.IdentificationPartRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/identification-parts", router.GetAll)
}
