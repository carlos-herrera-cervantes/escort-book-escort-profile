package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapPhotoRoutes(v *echo.Group) {
	router := &controllers.PhotoController{
		Repository: &repositories.PhotoRepository{
			Data: singleton.NewPostgresClient(),
		},
		S3Service: &services.S3Service{
			S3Client: singleton.NewS3Client(),
		},
	}

	v.GET("/escort/profile/photos", router.GetAll)
	v.GET("/escort/:id/profile/photos", router.GetById)
	v.POST("/escort/profile/photos", router.Create)
	v.DELETE("/escort/profile/photos/:id", router.DeleteOne)
}
