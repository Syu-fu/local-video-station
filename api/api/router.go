package api

import (
	"database/sql"

	"api/controllers"
	"api/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minio/minio-go/v7"
)

func NewRouter(db *sql.DB, mc *minio.Client) *echo.Echo {
	ser := services.NewMyAppService(db, mc)
	aCon := controllers.NewVideoController(ser)
	tCon := controllers.NewTagController(ser)
	e := echo.New()

	e.GET("/video/list", aCon.VideoListHandler)
	e.GET("/video/count", aCon.VideoCountHandler)
	e.GET("/video/:id", aCon.VideoDetailHandler)
	e.POST("/video", aCon.PostVideoHandler)
	e.PUT("/video", aCon.PutVideoHandler)

	e.GET("/tag/list", tCon.TagListHandler)
	e.POST("/tag", tCon.PostTagHandler)

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	return e
}
