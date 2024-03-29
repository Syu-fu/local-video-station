package controllers

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"api/apperrors"
	"api/controllers/services"
	"api/models"

	"github.com/google/uuid"
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

	tagIDs := ctx.QueryParam("tags")

	var videoList []models.Video

	if tagIDs != "" {
		videoList, err = c.service.GetVideoListByTagsService(tagIDs, page)
	} else {
		videoList, err = c.service.GetVideoListService(page)
	}

	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, videoList)
}

func (c *VideoController) VideoCountHandler(ctx echo.Context) error {
	tagIDs := ctx.QueryParam("tags")

	var count int

	var err error

	if tagIDs != "" {
		count, err = c.service.GetVideoCountByTagsService(tagIDs)
	} else {
		count, err = c.service.GetVideoCountService()
	}

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

func (c *VideoController) VideoDetailHandler(ctx echo.Context) error {
	videoID := ctx.Param("id")

	video, err := c.service.GetVideoService(videoID)
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, video)
}

func (c *VideoController) PostVideoHandler(ctx echo.Context) error {
	thumbnailFile, thumbnailHeader, err := ctx.Request().FormFile("thumbnail")
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}
	defer thumbnailFile.Close()

	thumbnailReader := bufio.NewReader(thumbnailFile)

	videoFile, videoHeader, err := ctx.Request().FormFile("video")
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}
	defer videoFile.Close()

	videoReader := bufio.NewReader(videoFile)

	video := new(models.Video)

	thumbnailPart := strings.Split(thumbnailHeader.Filename, ".")
	thumbnailExtension := "." + thumbnailPart[len(thumbnailPart)-1]
	videoPart := strings.Split(videoHeader.Filename, ".")
	videoExtension := "." + videoPart[len(videoPart)-1]

	video.ID = uuid.New().String()
	video.Title = ctx.Request().FormValue("title")
	video.TitleReading = ctx.Request().FormValue("titleReading")
	video.Url = "video/" + video.ID + videoExtension
	video.ThumbnailUrl = "thumbnail/" + video.ID + thumbnailExtension
	tagsStr := ctx.Request().FormValue("tags")

	var tags []models.Tag
	if err := json.Unmarshal([]byte(tagsStr), &tags); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")

		return apperrors.ErrorHandler(ctx, err)
	}

	video.Tags = tags

	response, err := c.service.PostVideoService(*video, thumbnailReader, videoReader)
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *VideoController) PutVideoHandler(ctx echo.Context) error {
	thumbnailFile, _, err := ctx.Request().FormFile("thumbnail")

	var thumbnailReader *bufio.Reader

	if err == nil {
		defer thumbnailFile.Close()

		thumbnailReader = bufio.NewReader(thumbnailFile)
	}

	videoFile, _, err := ctx.Request().FormFile("video")

	var videoReader *bufio.Reader

	if err == nil {
		defer videoFile.Close()

		videoReader = bufio.NewReader(videoFile)
	}

	video := new(models.Video)

	video.ID = ctx.Request().FormValue("id")
	video.Title = ctx.Request().FormValue("title")
	video.TitleReading = ctx.Request().FormValue("titleReading")
	tagsStr := ctx.Request().FormValue("tags")

	var tags []models.Tag
	if err := json.Unmarshal([]byte(tagsStr), &tags); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")

		return apperrors.ErrorHandler(ctx, err)
	}

	video.Tags = tags

	response, err := c.service.PutVideoService(*video, thumbnailReader, videoReader)
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response)
}
