package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapIdentificationPartRoutes(v *echo.Group) {
	router := &controllers.IdentificationPartController{
		Repository: &repositories.IdentificationPartRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/identification-parts", router.GetAll)
}
