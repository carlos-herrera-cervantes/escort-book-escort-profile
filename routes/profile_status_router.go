package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"

	"github.com/labstack/echo/v4"
)

func BoostrapProfileStatusRoutes(v *echo.Group) {
	router := &controllers.ProfileStatusController{
		Repository: &repositories.ProfileStatusRepository{
			Data: db.New(),
		},
		ProfileStatusCategoryRepository: &repositories.ProfileStatusCategoryRepository{
			Data: db.New(),
		},
		KafkaService: &services.KafkaService{
			Producer: db.NewProducer(),
		},
	}

	v.GET("/escort/:id/profile/status", router.GetByExternal)
	v.PATCH("/escort/:id/profile/status", router.UpdateByExternal)
	v.PATCH("/escort/profile/status", router.UpdateOne)
}
