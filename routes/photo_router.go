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

	v.GET("/escort/profile/photos", router.GetAll)
	v.GET("/escort/:id/profile/photos", router.GetById)
	v.POST("/escort/profile/photos", router.Create)
	v.DELETE("/escort/profile/photos/:id", router.DeleteOne)
}
