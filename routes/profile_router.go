package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapProfileRoutes(v *echo.Group) {
	router := &controllers.ProfileController{
		Repository: &repositories.ProfileRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/profile/:id", router.GetOne)
}
