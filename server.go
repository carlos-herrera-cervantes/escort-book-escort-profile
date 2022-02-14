package main

import (
	"escort-book-escort-profile/listeners"
	"escort-book-escort-profile/routes"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	v1 := e.Group("/api/v1")

	listeners.BoostrapListeners()

	routes.BoostrapPhotoRoutes(v1)
	routes.BoostrapIdentificationRoutes(v1)
	routes.BoostrapIdentificationPartRoutes(v1)
	routes.BoostrapAvatarRoutes(v1)
	routes.BoostrapScheduleRoutes(v1)
	routes.BoostrapTimeCategoryRoutes(v1)
	routes.BoostrapPriceRoutes(v1)
	routes.BoostrapNationalityRoutes(v1)
	routes.BoostrapDayRoutes(v1)
	routes.BoostrapAttentionSiteCategoryRoutes(v1)
	routes.BoostrapAttentionSiteRoutes(v1)
	routes.BoostrapBiographyRoutes(v1)
	routes.BoostrapProfileStatusCategoryRoutes(v1)
	routes.BoostrapProfileStatusRoutes(v1)
	routes.BoostrapProfileRoutes(v1)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
