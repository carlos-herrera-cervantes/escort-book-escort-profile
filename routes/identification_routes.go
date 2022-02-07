package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"

	"github.com/labstack/echo/v4"
)

func BoostrapIdentificationRoutes(v *echo.Group) {
	router := &controllers.IdentificationController{
		Repository: &repositories.IdentificationRepository{
			Data: db.New(),
		},
		S3Service: &services.S3Service{},
	}

	v.GET("/escort/profile/:profileId/identifications", router.GetAll)
	v.POST("/escort/profile/:profileId/identifications", router.Create)
	v.PATCH("/escort/profile/:profileId/identifications/:id", router.UpdateOne)
}
