package services

import (
	"api/apperrors"
	"api/models"
	"api/repositories"
)

func (s *MyAppService) PostTagService(tag models.Tag) (models.Tag, error) {
	newTag, err := repositories.InsertTag(s.db, tag)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")

		return models.Tag{}, err
	}

	return newTag, nil
}

func (s *MyAppService) GetTagListService() ([]models.Tag, error) {
	tagList, err := repositories.SelectAllTags(s.db)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")

		return nil, err
	}

	if len(tagList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")

		return nil, err
	}

	return tagList, nil
}
