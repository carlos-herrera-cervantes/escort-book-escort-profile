package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapProfileStatusRoutes(v *echo.Group) {
	router := &controllers.ProfileStatusController{
		Repository: &repositories.ProfileStatusRepository{
			Data: singleton.NewPostgresClient(),
		},
		ProfileStatusCategoryRepository: &repositories.ProfileStatusCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
		KafkaService: &services.KafkaService{
			Producer: singleton.NewProducer(),
		},
	}

	v.GET("/escort/:id/profile/status", router.GetByExternal)
	v.PATCH("/escort/:id/profile/status", router.UpdateByExternal)
	v.PATCH("/escort/profile/status", router.UpdateOne)
}
