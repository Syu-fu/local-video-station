package services

import "api/models"

type VideoServicer interface {
	GetVideoListService(page int) ([]models.Video, error)
	GetVideoCountService() (int, error)
}
