package routes

import (
	"escort-book-escort-profile/controllers"
	"escort-book-escort-profile/db"
	"escort-book-escort-profile/repositories"

	"github.com/labstack/echo/v4"
)

func BoostrapAttentionSiteCategoryRoutes(v *echo.Group) {
	router := &controllers.AttentionSiteCategoryController{
		Repository: &repositories.AttentionSiteCategoryRepository{
			Data: db.New(),
		},
	}

	v.GET("/escort/attention-site-categories", router.GetAll)
}
