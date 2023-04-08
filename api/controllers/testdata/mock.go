package testdata

import "api/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) GetVideoListService(page int) ([]models.Video, error) {
	return VideoTestData, nil
}

func (s *serviceMock) GetVideoCountService() (int, error) {
	videoCount := 2

	return videoCount, nil
}
