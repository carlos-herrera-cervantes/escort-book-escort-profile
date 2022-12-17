package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapIdentificationRoutes(v *echo.Group) {
	router := &controllers.IdentificationController{
		Repository: &repositories.IdentificationRepository{
			Data: singleton.NewPostgresClient(),
		},
		S3Service: &services.S3Service{
			S3Client: singleton.NewS3Client(),
		},
		IdentificationCategoryRepository: &repositories.IdentificationPartRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/:id/profile/identifications", router.GetByExternal)
	v.GET("/escort/profile/identifications", router.GetAll)
	v.POST("/escort/profile/identifications", router.Create)
	v.PATCH("/escort/profile/identifications/:id", router.UpdateOne)
}
