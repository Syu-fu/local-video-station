package api

import (
	"database/sql"

	"api/controllers"
	"api/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(db *sql.DB) *echo.Echo {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewVideoController(ser)
	tCon := controllers.NewTagController(ser)
	e := echo.New()

	e.GET("/video/list", aCon.VideoListHandler)
	e.GET("/video/count", aCon.VideoCountHandler)
	e.GET("/video/:id", aCon.VideoDetailHandler)

	e.POST("/tag", tCon.PostTagHandler)

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	return e
}
