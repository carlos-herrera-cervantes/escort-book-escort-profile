package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"

	"github.com/labstack/echo/v4"
)

func BoostrapPhotoRoutes(v *echo.Group) {
	router := &controllers.PhotoController{
		Repository: &repositories.PhotoRepository{
			Data: db.New(),
		},
		S3Service: &services.S3Service{},
	}

	v.GET("/escort/profile/:profileId/photos", router.GetAll)
	v.POST("/escort/profile/:profileId/photos", router.Create)
	v.DELETE("/escort/profile/:profileId/photos/:id", router.DeleteOne)
}
