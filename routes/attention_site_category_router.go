package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapAttentionSiteCategoryRoutes(v *echo.Group) {
	router := &controllers.AttentionSiteCategoryController{
		Repository: &repositories.AttentionSiteCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/attention-site-categories", router.GetAll)
}
