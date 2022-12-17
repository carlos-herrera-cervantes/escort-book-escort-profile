package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapProfileRoutes(v *echo.Group) {
	router := &controllers.ProfileController{
		Repository: &repositories.ProfileRepository{
			Data: singleton.NewPostgresClient(),
		},
		Emitter: &services.EmitterService{
			Emitter: singleton.NewEmitter(),
		},
		NationalityRepository: &repositories.NationalityRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escorts", router.GetAll)
	v.GET("/escort/profile", router.GetOne)
	v.GET("/escort/:id/profile", router.GetById)
	v.POST("/escort/profile", router.Create)
	v.PATCH("/escort/profile", router.UpdateOne)
	v.DELETE("/escort/profile", router.DeleteOne)
}
