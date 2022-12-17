package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapNationalityRoutes(v *echo.Group) {
	router := &controllers.NationalityController{
		Repository: &repositories.NationalityRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/nationalities", router.GetAll)
}
