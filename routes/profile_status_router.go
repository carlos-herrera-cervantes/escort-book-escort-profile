package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

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
	}

	v.PATCH("/escort/profile/status", router.UpdateOne)
}
