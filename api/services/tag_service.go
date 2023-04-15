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
