package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapAttentionSiteRoutes(v *echo.Group) {
	router := &controllers.AttentionSiteController{
		Repository: &repositories.AttentionSiteRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/profile/attention-sites", router.GetAll)
	v.GET("/escort/profile/attention-sites/:id", router.GetById)
	v.POST("/escort/profile/attention-sites", router.Create)
	v.DELETE("/escort/profile/attention-sites/:id", router.DeleteOne)
}
