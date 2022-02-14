package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapNationalityRoutes(v *echo.Group) {
	router := &controllers.NationalityController{
		Repository: &repositories.NationalityRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/nationalities", router.GetAll)
}
