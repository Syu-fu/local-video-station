package services

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"sync"

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

func (s *MyAppService) GetVideoCountService() (int, error) {
	videoCount, err := repositories.SelectVideoCount(s.db)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")

		return 0, err
	}

	return videoCount, nil
}

func (s *MyAppService) GetVideoService(videoID string) (models.Video, error) {
	var video models.Video

	var tagList []models.Tag

	var videoGetErr, tagGetErr error

	var amu sync.Mutex

	var cmu sync.Mutex

	var wg sync.WaitGroup

	waitCount := 2
	wg.Add(waitCount)

	go func(db *sql.DB, videoID string) {
		defer wg.Done()
		amu.Lock()
		video, videoGetErr = repositories.SelectVideoDetail(db, videoID)
		amu.Unlock()
	}(s.db, videoID)

	go func(db *sql.DB, videoID string) {
		defer wg.Done()
		cmu.Lock()
		tagList, tagGetErr = repositories.SelectTagList(db, videoID)
		cmu.Unlock()
	}(s.db, videoID)

	wg.Wait()

	if videoGetErr != nil {
		if errors.Is(videoGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(videoGetErr, "no data")

			return models.Video{}, err
		}

		err := apperrors.GetDataFailed.Wrap(videoGetErr, "fail to get data")

		return models.Video{}, err
	}

	if tagGetErr != nil {
		tagGetErr = apperrors.GetDataFailed.Wrap(tagGetErr, "fail to get data")

		return models.Video{}, tagGetErr
	}

	video.Tags = tagList

	return video, nil
}

func (s *MyAppService) PostVideoService(video models.Video, thumbnailFile io.Reader, videoFile io.Reader) (models.Video, error) {
	newVideo, err := repositories.InsertVideo(s.db, video)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")

		return models.Video{}, err
	}

	_, err = repositories.InsertVideoTag(s.db, video)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")

		return models.Video{}, err
	}

	ctx := context.Background()
	bucketName := "data"

	if err := repositories.SaveFile(s.mc, ctx, bucketName, video.ThumbnailUrl, thumbnailFile, -1, "image/jpeg"); err != nil {
		return models.Video{}, err
	}

	if err := repositories.SaveFile(s.mc, ctx, bucketName, video.Url, videoFile, -1, "video/mp4"); err != nil {
		return models.Video{}, err
	}

	return newVideo, nil
}

func (s *MyAppService) GetVideoListByTagsService(tagIDs string, page int) ([]models.Video, error) {
	videoList, err := repositories.SelectVideoListByTags(s.db, tagIDs, page)
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
