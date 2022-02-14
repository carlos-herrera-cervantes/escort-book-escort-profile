package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapProfileStatusCategoryRoutes(v *echo.Group) {
	router := &controllers.ProfileStatusCategoryController{
		Repository: &repositories.ProfileStatusCategoryRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/profile-status-categories", router.GetAll)
}
