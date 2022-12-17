package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapProfileStatusCategoryRoutes(v *echo.Group) {
	router := &controllers.ProfileStatusCategoryController{
		Repository: &repositories.ProfileStatusCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/profile-status-categories", router.GetAll)
}
