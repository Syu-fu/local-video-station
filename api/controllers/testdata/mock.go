package testdata

import (
	"io"

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

func (s *serviceMock) PostVideoService(video models.Video, thumbnailFile io.Reader, videoFile io.Reader) (models.Video, error) {
	return models.Video{}, nil
}

func (s *serviceMock) GetVideoListByTagsService(tagIDs string, page int) ([]models.Video, error) {
	return VideoTestData, nil
}

func (s *serviceMock) GetVideoCountByTagsService(tagIDs string) (int, error) {
	return 0, nil
}
