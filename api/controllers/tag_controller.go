package controllers

import (
	"net/http"

	"api/apperrors"
	"api/controllers/services"
	"api/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TagController struct {
	service services.TagServicer
}

func NewTagController(s services.TagServicer) *TagController {
	return &TagController{service: s}
}

func (c *TagController) PostTagHandler(ctx echo.Context) error {
	tag := new(models.Tag)
	if err := ctx.Bind(tag); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")

		return apperrors.ErrorHandler(ctx, err)
	}

	tag.ID = uuid.New().String()

	newTag, err := c.service.PostTagService(*tag)
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, newTag)
}

func (c *TagController) TagListHandler(ctx echo.Context) error {
	tagList, err := c.service.GetTagListService()
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, tagList)
}
