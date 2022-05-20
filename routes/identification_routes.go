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
		IdentificationCategoryRepository: &repositories.IdentificationPartRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/:id/profile/identifications", router.GetByExternal)
	v.GET("/escort/profile/identifications", router.GetAll)
	v.POST("/escort/profile/identifications", router.Create)
	v.PATCH("/escort/profile/identifications/:id", router.UpdateOne)
}
