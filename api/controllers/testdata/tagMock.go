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
