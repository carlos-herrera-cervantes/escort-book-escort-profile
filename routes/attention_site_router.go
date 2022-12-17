package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/repositories"
	"escort-book-escort-profile/singleton"

	"github.com/labstack/echo/v4"
)

func BootstrapAttentionSiteRoutes(v *echo.Group) {
	router := &controllers.AttentionSiteController{
		Repository: &repositories.AttentionSiteRepository{
			Data: singleton.NewPostgresClient(),
		},
		AttentionSiteCategoryRepository: &repositories.AttentionSiteCategoryRepository{
			Data: singleton.NewPostgresClient(),
		},
	}

	v.GET("/escort/profile/attention-sites", router.GetAll)
	v.GET("/escort/:id/profile/attention-sites", router.GetById)
	v.POST("/escort/profile/attention-sites", router.Create)
	v.DELETE("/escort/profile/attention-sites/:id", router.DeleteOne)
}
