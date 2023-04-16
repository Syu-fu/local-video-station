package services

import (
	"io"

	"api/models"
)

type VideoServicer interface {
	GetVideoListService(page int) ([]models.Video, error)
	GetVideoCountService() (int, error)
	GetVideoService(videoID string) (models.Video, error)
	PostVideoService(video models.Video, thumbnailFile io.Reader, videoFile io.Reader) (models.Video, error)
}

type TagServicer interface {
	PostTagService(tag models.Tag) (models.Tag, error)
	GetTagListService() ([]models.Tag, error)
}
