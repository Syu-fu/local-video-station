package controllers

import (
	"net/http"
	"strconv"

	"api/apperrors"
	"api/controllers/services"

	"github.com/labstack/echo/v4"
)

type VideoController struct {
	service services.VideoServicer
}

func NewVideoController(s services.VideoServicer) *VideoController {
	return &VideoController{service: s}
}

func (c *VideoController) VideoListHandler(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")

		return apperrors.ErrorHandler(ctx, err)
	}

	videoList, err := c.service.GetVideoListService(page)
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, videoList)
}

func (c *VideoController) VideoCountHandler(ctx echo.Context) error {
	count, err := c.service.GetVideoCountService()
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	response := struct {
		Count int `json:"count"`
	}{
		Count: count,
	}

	return ctx.JSON(http.StatusOK, response)
}
