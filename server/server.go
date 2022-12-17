package server

import (
	"fmt"

	"escort-book-escort-profile/config"
	"escort-book-escort-profile/listeners"
	"escort-book-escort-profile/routes"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	v1 := e.Group("/api/v1")

	listeners.BootstrapListeners()

	routes.BootstrapPhotoRoutes(v1)
	routes.BootstrapIdentificationRoutes(v1)
	routes.BootstrapIdentificationPartRoutes(v1)
	routes.BootstrapAvatarRoutes(v1)
	routes.BootstrapScheduleRoutes(v1)
	routes.BootstrapTimeCategoryRoutes(v1)
	routes.BootstrapPriceRoutes(v1)
	routes.BootstrapNationalityRoutes(v1)
	routes.BootstrapDayRoutes(v1)
	routes.BootstrapAttentionSiteCategoryRoutes(v1)
	routes.BootstrapAttentionSiteRoutes(v1)
	routes.BootstrapBiographyRoutes(v1)
	routes.BootstrapProfileStatusCategoryRoutes(v1)
	routes.BootstrapProfileStatusRoutes(v1)
	routes.BootstrapProfileRoutes(v1)
	routes.BootstrapServiceRoutes(v1)
	routes.BootstrapServiceCategoryRoutes(v1)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.InitApp().Port)))
}
