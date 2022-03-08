package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"

	"github.com/labstack/echo/v4"
)

func BoostrapProfileRoutes(v *echo.Group) {
	router := &controllers.ProfileController{
		Repository: &repositories.ProfileRepository{
			Data: db.New(),
		},
		Emitter: &services.EmitterService{},
	}

	v.GET("/escort/profile", router.GetOne)
	v.POST("/escort/profile", router.Create)
	v.PATCH("/escort/profile", router.UpdateOne)
	v.DELETE("/escort/profile", router.DeleteOne)
}
