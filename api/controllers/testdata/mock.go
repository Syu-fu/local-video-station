package testdata

import (
	"api/apperrors"
	"api/models"
)

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

func (s *serviceMock) GetVideoService(id string) (models.Video, error) {
	if id == "3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4" {
		return VideoTestData[0], nil
	}

	var err error
	err = apperrors.NAData.Wrap(err, "no data")

	return models.Video{}, err
}
