package services

import (
	"api/apperrors"
	"api/models"
	"api/repositories"
)

func (s *MyAppService) GetVideoListService(page int) ([]models.Video, error) {
	videoList, err := repositories.SelectVideoList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")

		return nil, err
	}

	if len(videoList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")

		return nil, err
	}

	return videoList, nil
}
