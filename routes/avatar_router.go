package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapAvatarRoutes(v *echo.Group) {
	router := &controllers.AvatarController{
		Repository: &repositories.AvatarRepository{
			Data: singleton.NewPostgresClient(),
		},
		S3Service: &services.S3Service{
			S3Client: singleton.NewS3Client(),
		},
	}

	v.GET("/escort/profile/avatar", router.GetOne)
	v.GET("/escort/:id/profile/avatar", router.GetById)
	v.PATCH("/escort/profile/avatar", router.Upsert)
	v.DELETE("/escort/profile/avatar", router.DeleteOne)
}
