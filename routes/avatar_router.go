package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/services"

	"github.com/labstack/echo/v4"
)

func BoostrapAvatarRoutes(v *echo.Group) {
	router := &controllers.AvatarController{
		Repository: &repositories.AvatarRepository{
			Data: db.New(),
		},
		S3Service: &services.S3Service{},
	}

	v.GET("/escort/profile/:profileId/avatar", router.GetOne)
	v.POST("/escort/profile/:profileId/avatar", router.Create)
	v.PATCH("/escort/profile/:profileId/avatar", router.UpdateOne)
	v.DELETE("/escort/profile/:profileId/avatar", router.DeleteOne)
}
