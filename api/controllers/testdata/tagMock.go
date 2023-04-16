package testdata

import (
	"api/models"
)

type tagServiceMock struct{}

func NewTagServiceMock() *tagServiceMock {
	return &tagServiceMock{}
}

func (s *tagServiceMock) PostTagService(tag models.Tag) (models.Tag, error) {
	return tag, nil
}

func (s *tagServiceMock) GetTagListService() ([]models.Tag, error) {
	return []models.Tag{}, nil
}
